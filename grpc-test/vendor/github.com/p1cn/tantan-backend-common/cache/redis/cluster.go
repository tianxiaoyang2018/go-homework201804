package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/p1cn/tantan-backend-common/config"
)

type clusterClient struct {
	clusterName string

	cluster *redis.ClusterClient
	config  config.RedisConfig
}

var clusterId int = 1

func NewClusterClient(redisConfig config.RedisConfig) (*clusterClient, error) {
	cc := &clusterClient{
		config: redisConfig,
	}

	if redisConfig.ClusterName != "" {
		cc.clusterName = redisConfig.ClusterName
	} else {
		cc.clusterName = fmt.Sprintf("cluster%d", clusterId)
		clusterId++
	}

	clusterAddrs := make([]string, 0, len(redisConfig.Addrs))

	for _, redisAddr := range redisConfig.Addrs {
		clusterAddrs = append(clusterAddrs, redisAddr.Addr)
	}

	option := redis.ClusterOptions{
		Addrs:        clusterAddrs,
		PoolTimeout:  time.Duration(redisConfig.PoolTimeout),
		PoolSize:     redisConfig.PoolSize,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout),
		WriteTimeout: time.Duration(redisConfig.WriteTimeout),
	}
	cc.cluster = redis.NewClusterClient(&option)
	if err := cc.cluster.Ping().Err(); err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *clusterClient) Get(ctx context.Context, key string) (string, error) {
	record := ic.timer.Timer()
	cmd := cc.cluster.Get(key)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			record(cc.clusterName, "get", MISS)
		} else {
			record(cc.clusterName, "get", ERR)
		}
	} else {
		record(cc.clusterName, "get", OK)
	}
	return cmd.Result()
}

func (cc *clusterClient) Exist(ctx context.Context, key string) (bool, error) {
	record := ic.timer.Timer()
	cmd := cc.cluster.Exists(key)
	if cmd.Err() != nil {
		record(cc.clusterName, "exist", ERR)
		return false, cmd.Err()
	}
	record(cc.clusterName, "exist", OK)
	if cmd.Val() == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (cc *clusterClient) Gets(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) <= 0 {
		return nil, errors.New("keys is empty")
	}
	record := ic.timer.Timer()
	pipe := cc.cluster.Pipeline()
	defer pipe.Close()
	pipelineCmds := make([]*redis.StringCmd, 0, len(keys))
	for _, key := range keys {
		pipelineCmds = append(pipelineCmds, pipe.Get(key))
	}
	_, err := pipe.Exec()
	if err != nil && err != redis.Nil {
		record(cc.clusterName, "gets", ERR)
		return nil, err
	}

	res := make([]interface{}, 0, len(keys))
	for _, pcmd := range pipelineCmds {
		value, err := pcmd.Result()
		if err != nil && err != redis.Nil {
			record(cc.clusterName, "gets", ERR)
			return nil, err
		}
		if err == redis.Nil {
			record(cc.clusterName, "gets", MISS)
			res = append(res, Nil)
		} else {
			record(cc.clusterName, "gets", OK)
			res = append(res, value)
		}
	}

	return res, nil
}

func (cc *clusterClient) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	var ex time.Duration
	if len(expiration) > 0 {
		ex = expiration[0]
	} else {
		ex = 0
	}
	record := ic.timer.Timer()
	status := cc.cluster.Set(key, value, ex)
	if status.Err() != nil {
		record(cc.clusterName, "set", ERR)
		return status.Err()
	}
	record(cc.clusterName, "set", OK)
	return nil
}

func (cc *clusterClient) Sets(ctx context.Context, kvs map[string]interface{}, expiration ...time.Duration) error {
	if len(kvs) <= 0 {
		return errors.New("kvs is empty")
	}

	var ex time.Duration
	if len(expiration) > 0 {
		ex = expiration[0]
	} else {
		ex = 0
	}

	record := ic.timer.Timer()
	pipe := cc.cluster.Pipeline()
	defer pipe.Close()
	pipelineCmds := make([]*redis.StatusCmd, 0, len(kvs))
	for key, value := range kvs {
		pipelineCmds = append(pipelineCmds, pipe.Set(key, value, ex))
	}
	_, err := pipe.Exec()
	if err != nil {
		record(cc.clusterName, "sets", ERR)
		return err
	}

	for _, pcmd := range pipelineCmds {
		_, err := pcmd.Result()
		if err != nil {
			record(cc.clusterName, "sets", ERR)
			return err
		}
	}

	record(cc.clusterName, "sets", OK)
	return nil
}

func (cc *clusterClient) Del(ctx context.Context, keys []string) error {
	record := ic.timer.Timer()
	status := cc.cluster.Del(keys...)
	if status.Err() != nil {
		record(cc.clusterName, "del", ERR)
		return status.Err()
	}
	record(cc.clusterName, "del", OK)
	return nil
}

func (cc *clusterClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	record := ic.timer.Timer()
	status := cc.cluster.Expire(key, expiration)
	if status.Err() != nil {
		record(cc.clusterName, "expire", ERR)
		return status.Err()
	}
	record(cc.clusterName, "expire", OK)
	return nil
}

func (cc *clusterClient) Ping(ctx context.Context) error {
	record := ic.timer.Timer()
	err := cc.cluster.Ping().Err()
	if err == nil {
		record(cc.clusterName, "ping", OK)
	} else {
		record(cc.clusterName, "ping", ERR)
	}
	return err
}
