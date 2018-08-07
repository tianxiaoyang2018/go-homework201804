package postgres_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/db/postgres"
	"github.com/p1cn/tantan-backend-common/health"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	pg "gopkg.in/pg.v3"
)

type TestConfig struct {
	Postgres map[string]config.PostgresCluster
}

var tConfig TestConfig

var (
	configFile = flag.String("config", "", "config file path")
	dbName     = flag.String("dbname", "", "db name")
	metric     = flag.Bool("metric", false, "print metrics")
)

func initConfig(t *testing.T) {
	flag.Parse()

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(data, &tConfig)
	if err != nil {
		t.Error(err)
	}

	log.Init(log.Config{
		Output: []string{"stderr", "syslog"},
		Level:  "debug",
		Flags:  []string{"file", "level", "date"},
	})

	health.Init(":9090")

}

func TestGet(t *testing.T) {
	initConfig(t)

	if len(tConfig.Postgres) == 0 {
		data, _ := json.Marshal(tConfig)
		t.Error(string(data))
	}

	manager, err := postgres.NewDB(postgres.Config{
		Name:     "test",
		Cluster:  tConfig.Postgres[*dbName],
		Graceful: nil,
	})

	if err != nil {
		t.Error(err)
	}
	db := manager.Get(postgres.DbWrite)

	var res pgRoles
	_, err = db.Query(&res, "select rolname from pg_roles;")
	if err != nil {
		t.Error(err)
	}

	// read
	db = manager.Get(postgres.DbRead)

	_, err = db.Query(&res, "select rolname from pg_roles;")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Roles)

	// test prometheus
	_, _ = db.Query(&res, "select rolname from pg_roles()")

	if *metric {
		fmt.Println(metrics.GetPromethuesAsFmtText())
	}
}

func TestGetShard(t *testing.T) {
	initConfig(t)

	if len(tConfig.Postgres) == 0 {
		data, _ := json.Marshal(tConfig)
		t.Error(string(data))
	}

	manager, err := postgres.NewDB(postgres.Config{
		Name:     "test",
		Cluster:  tConfig.Postgres[*dbName],
		Graceful: nil,
	})
	if err != nil {
		t.Error(err)
	}
	var res Messages

	db, shardNumber, err := manager.GetByNumber(postgres.DbRead, 0)
	if err != nil {
		t.Error(err)
	}

	q, err := manager.WithSchema(shardNumber, "select * from %v.messages limit 10")
	if err != nil {
		t.Error(err)
	}

	_, err = db.Query(&res, q)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.C)
}

func TestGetShardByNumber(t *testing.T) {
	initConfig(t)

	if len(tConfig.Postgres) == 0 {
		data, _ := json.Marshal(tConfig)
		t.Error(string(data))
	}

	manager, err := postgres.NewDB(postgres.Config{
		Name:     "test",
		Cluster:  tConfig.Postgres[*dbName],
		Graceful: nil,
	})
	if err != nil {
		t.Error(err)
	}
	var res Messages

	db, err := manager.GetByShardNumber(postgres.DbRead, 0)
	if err != nil {
		t.Error(err)
	}

	_, err = db.Query(&res, "select * from rel_8192_0.messages limit 10")
	if err != nil {
		t.Error(err)
	}

	t.Log(res.C)
}

func TestWalkShards(t *testing.T) {
	initConfig(t)

	if len(tConfig.Postgres) == 0 {
		data, _ := json.Marshal(tConfig)
		t.Error(string(data))
	}

	manager, err := postgres.NewDB(postgres.Config{
		Name:     "test",
		Cluster:  tConfig.Postgres[*dbName],
		Graceful: nil,
	})
	if err != nil {
		t.Error(err)
	}

	shardsCount := manager.GetShardCount()
	wait := sync.WaitGroup{}
	wait.Add(shardsCount)
	manager.WalkShards(func(db postgres.DBWrapper) {
		var res pgRoles
		ret, err := db.Query(&res, "select rolname from pg_roles;")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ret)
		wait.Done()
	})
	wait.Wait()
}

///////
type pgRole struct {
	Name string `pg:"rolname"`
}

type pgRoles struct {
	Roles []pgRole
}

func (self *pgRoles) NewRecord() interface{} {
	return &pgRole{}
}

type Message struct {
	Id          string        `pg:"id"`
	UserId      string        `pg:"user_id"`
	OtherUserId string        `pg:"other_user_id"`
	StickerId   string        `pg:"sticker_id"`
	QuestionId  string        `pg:"question_id"`
	ReferenceId string        `pg:"reference_id"`
	MomentId    string        `pg:"moment_id"`
	Value       string        `pg:"value"`
	Recalled    bool          `pg:"recalled"`
	CreatedTime postgres.Time `pg:"created_time"`
	UpdatedTime postgres.Time `pg:"updated_time"`
	SentFrom    string        `pg:"sent_from"`
	Location    string        `pg:"location"`
	Status      string        `pg:"status"`
}

type Messages struct {
	C []Message
}

var _ pg.Collection = &Messages{}

func (self *Messages) NewRecord() interface{} {
	self.C = append(self.C, Message{})
	return &self.C[len(self.C)-1]

}
