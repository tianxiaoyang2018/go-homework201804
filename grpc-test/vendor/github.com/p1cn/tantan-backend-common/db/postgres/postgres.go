package postgres

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/p1cn/tantan-backend-common/config"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-backend-common/util"
)

type dBCluster struct {
	mod               int
	shardMapByMod     map[int]*dbGroup
	shards            []*dbGroup
	roundRobin        uint64
	shardSchemaPrefix string
}

type dbGroup struct {
	physicalShardId int

	master   DBWrapper
	readOnly []DBWrapper

	roundRobin uint64
}

// NewDB的配置
type Config struct {
	// 数据库名字
	Name string
	// 数据库集群的配置
	Cluster config.PostgresCluster
	// 优雅退出接口
	Graceful util.GracefulMonitor
}

var (
	metricsTimer *metrics.Timer
	once         sync.Once
)

// NewDB 初始化一个db的DBManager
func NewDB(cfg Config) (DBManager, error) {
	once.Do(func() {
		metricsTimer = metrics.NewTimer(metrics.NameSpaceTantan, "db_pg", "postgres metrics", []string{"db_name", "op_name", "ret"})
	})

	db := &dBCluster{
		mod:               cfg.Cluster.Mod,
		shardMapByMod:     map[int]*dbGroup{},
		shardSchemaPrefix: cfg.Cluster.SchemaPrefix,
	}

	ic := &defaultInfoCollector{}
	ic.timer = metricsTimer

	for physicalShardId, v := range cfg.Cluster.Shards {
		group := setupDbGroup(cfg.Name, ic, physicalShardId, v.Master, v.Slaves, newDBWrapperWithCB)
		db.shards = append(db.shards, group)
		for i := v.FromLogicalShardMod; i <= v.ToLogicalShardMod; i++ {
			_, ok := db.shardMapByMod[i]
			if ok {
				return nil, fmt.Errorf("Logical shard %v configured twice", i)
			}
			db.shardMapByMod[i] = group
		}
	}

	slog.Info("initialized database successfully : name(%s)", cfg.Name)

	return db, nil
}

type infoCollector interface {
	Timer() func(values ...string)
}

type defaultInfoCollector struct {
	timer *metrics.Timer
}

func (t *defaultInfoCollector) Timer() func(values ...string) {
	if t == nil || t.timer == nil {
		return func(values ...string) {}
	}
	return t.timer.Timer()
}

func setupDbGroup(dbName string, ic infoCollector, physicalShardId int, dbConf config.PostgreSql, dbReadOnlyConf []config.PostgreSql, pgdbwInit func(string, infoCollector, config.PostgreSql) DBWrapper) *dbGroup {
	group := &dbGroup{physicalShardId: physicalShardId}
	if group.master == nil {
		group.master = pgdbwInit(dbName, ic, dbConf)
	}
	if group.readOnly == nil {
		group.readOnly = []DBWrapper{}
		for _, dbConfRo := range dbReadOnlyConf {
			group.readOnly = append(group.readOnly, pgdbwInit(dbName, ic, dbConfRo))
		}
	}
	return group
}

func (self *dBCluster) Get(op DbOperation) DBWrapper {
	// @todo opts
	rr := atomic.AddUint64(&self.roundRobin, 1)
	db := self.shards[rr%uint64(len(self.shards))]
	return self.getFromGroup(op, db)
}

//
func (self *dBCluster) GetByNumber(op DbOperation, id int) (DBWrapper, int, error) {

	group, ok := self.shardMapByMod[id%self.mod]
	if !ok {
		err := fmt.Errorf("Shard not properly setup for id(%v) , mod(%v)", id, self.mod)
		slog.Err("%v", err)
		return nil, -1, err
	}
	return self.getFromGroup(op, group), id % self.mod, nil
}

func (self *dBCluster) GetByShardNumber(op DbOperation, shardNum int) (DBWrapper, error) {
	group, ok := self.shardMapByMod[shardNum]
	if !ok {
		err := fmt.Errorf("Logical shard(%v) does not exist.", shardNum)
		slog.Err("%v", err)
		return nil, err
	}
	return self.getFromGroup(op, group), nil
}

func (self *dBCluster) WalkShards(f func(db DBWrapper)) {
	for _, v := range self.shards {
		f(v.master)
	}
}

func (self *dBCluster) getFromGroup(op DbOperation, group *dbGroup) DBWrapper {
	if op&DbWrite > 0 {
		return group.master
	}
	if op&DbRead > 0 && len(group.readOnly) > 0 {
		rr := atomic.AddUint64(&group.roundRobin, 1)
		slave := group.readOnly[rr%uint64(len(group.readOnly))]
		return slave
	}
	return group.master
}

func (self *dBCluster) MapUserIdsByShardId(userIds []string) (map[int][]string, error) {
	idsByShards := make(map[int][]string)
	for _, userId := range userIds {
		id, err := strconv.Atoi(userId)
		if err != nil {
			slog.Err("%v", err)
			return nil, err
		}
		shardId := util.Int32Abs(id) % self.mod
		idsByShards[shardId] = append(idsByShards[shardId], userId)
	}
	return idsByShards, nil
}

func (self *dBCluster) GetPhysicalShardNumberByLogicalShardNumber(shardNum int) (int, error) {
	group, ok := self.shardMapByMod[shardNum]
	if !ok {
		err := fmt.Errorf("Logical shard(%v) does not exist.", shardNum)
		slog.Err("%v", err)
		return -1, err
	}
	return group.physicalShardId, nil
}

func (self *dBCluster) WithSchema(shardNumber int, queryTemplate string) (string, error) {
	_, ok := self.shardMapByMod[shardNumber]
	if !ok {
		err := fmt.Errorf("Shard(%v) does not exist.", shardNumber)
		slog.Err("%v", err)
		return "", err
	}
	return fmt.Sprintf(queryTemplate, self.GetSchemaByShardNumber(shardNumber)), nil
}

func (self *dBCluster) GetSchemaByShardNumber(shardNum int) string {
	return self.schemaName(shardNum)
}

func (self *dBCluster) schemaName(shardNum int) string {
	prefix := strings.Trim(self.shardSchemaPrefix, "_")
	return fmt.Sprintf("%v_%v", prefix, shardNum)
}

func (self *dBCluster) GetShardCount() int {
	return len(self.shards)
}

func (self *dBCluster) GetLogicalShardCount() int {
	return self.mod
}
