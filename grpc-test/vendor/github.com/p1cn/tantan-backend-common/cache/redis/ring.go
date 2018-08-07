package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/p1cn/tantan-backend-common/config"
)

type ringClient struct {
	clusterName string

	ring   *redis.Ring
	config config.RedisConfig
}

var ringId int = 1

func NewRingClient(redisConfig config.RedisConfig) (*ringClient, error) {
	rc := &ringClient{
		config: redisConfig,
	}

	if redisConfig.ClusterName != "" {
		rc.clusterName = redisConfig.ClusterName
	} else {
		rc.clusterName = fmt.Sprintf("ring%d", ringId)
		ringId++
	}

	ringAddrs := make(map[string]string)

	for _, redisAddr := range redisConfig.Addrs {
		ringAddrs[redisAddr.Name] = redisAddr.Addr
	}

	option := redis.RingOptions{
		Addrs:              ringAddrs,
		HeartbeatFrequency: time.Duration(redisConfig.HeartbeatFrequency),
		DB:                 redisConfig.DB,
		Password:           redisConfig.Password,
		MaxRetries:         redisConfig.MaxRetries,

		DialTimeout:  time.Duration(redisConfig.DialTimeout),
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout),
		WriteTimeout: time.Duration(redisConfig.WriteTimeout),

		PoolTimeout:        time.Duration(redisConfig.PoolTimeout),
		PoolSize:           redisConfig.PoolSize,
		IdleTimeout:        time.Duration(redisConfig.IdleTimeout),
		IdleCheckFrequency: time.Duration(redisConfig.IdleCheckFrequency),
	}
	rc.ring = redis.NewRing(&option)
	if err := rc.ring.Ping().Err(); err != nil {
		return nil, err
	}
	return rc, nil
}

func (rc *ringClient) Get(ctx context.Context, key string) (string, error) {
	record := ic.timer.Timer()
	cmd := rc.ring.Get(key)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			record(rc.clusterName, "get", MISS)
		} else {
			record(rc.clusterName, "get", ERR)
		}
	} else {
		record(rc.clusterName, "get", OK)
	}
	return cmd.Result()
}

func (rc *ringClient) Exist(ctx context.Context, key string) (bool, error) {
	record := ic.timer.Timer()
	cmd := rc.ring.Exists(key)
	if cmd.Err() != nil {
		record(rc.clusterName, "exist", ERR)
		return false, cmd.Err()
	}
	record(rc.clusterName, "exist", OK)
	if cmd.Val() == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

//TODO - Need optimize: use multiple mget
func (rc *ringClient) Gets(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) <= 0 {
		return nil, errors.New("keys is empty")
	}
	record := ic.timer.Timer()
	pipe := rc.ring.Pipeline()
	defer pipe.Close()
	pipelineCmds := make([]*redis.StringCmd, 0, len(keys))
	for _, key := range keys {
		pipelineCmds = append(pipelineCmds, pipe.Get(key))
	}
	_, err := pipe.Exec()
	if err != nil && err != redis.Nil {
		record(rc.clusterName, "gets", ERR)
		return nil, err
	}

	res := make([]interface{}, 0, len(keys))
	for _, pcmd := range pipelineCmds {
		value, err := pcmd.Result()
		if err != nil && err != redis.Nil {
			record(rc.clusterName, "gets", ERR)
			return nil, err
		}
		if err == redis.Nil {
			record(rc.clusterName, "gets", MISS)
			res = append(res, Nil)
		} else {
			record(rc.clusterName, "gets", OK)
			res = append(res, value)
		}
	}

	return res, nil
}

func (rc *ringClient) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	var ex time.Duration
	if len(expiration) > 0 {
		ex = expiration[0]
	} else {
		ex = 0
	}
	record := ic.timer.Timer()
	status := rc.ring.Set(key, value, ex)
	if status.Err() != nil {
		record(rc.clusterName, "set", ERR)
		return status.Err()
	}
	record(rc.clusterName, "set", OK)
	return nil
}

func (rc *ringClient) Sets(ctx context.Context, kvs map[string]interface{}, expiration ...time.Duration) error {
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
	pipe := rc.ring.Pipeline()
	defer pipe.Close()
	pipelineCmds := make([]*redis.StatusCmd, 0, len(kvs))
	for key, value := range kvs {
		pipelineCmds = append(pipelineCmds, pipe.Set(key, value, ex))
	}
	_, err := pipe.Exec()
	if err != nil {
		record(rc.clusterName, "sets", ERR)
		return err
	}

	for _, pcmd := range pipelineCmds {
		_, err := pcmd.Result()
		if err != nil {
			record(rc.clusterName, "sets", ERR)
			return err
		}
	}

	record(rc.clusterName, "sets", OK)
	return nil
}

func (rc *ringClient) Del(ctx context.Context, keys []string) error {
	record := ic.timer.Timer()
	status := rc.ring.Del(keys...)
	if status.Err() != nil {
		record(rc.clusterName, "del", ERR)
		return status.Err()
	}
	record(rc.clusterName, "del", OK)
	return nil
}
func (rc *ringClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	record := ic.timer.Timer()
	status := rc.ring.Expire(key, expiration)
	if status.Err() != nil {
		record(rc.clusterName, "expire", ERR)
		return status.Err()
	}
	record(rc.clusterName, "expire", OK)
	return nil
}

func (rc *ringClient) Ping(ctx context.Context) error {
	record := ic.timer.Timer()
	err := rc.ring.Ping().Err()
	if err == nil {
		record(rc.clusterName, "ping", OK)
	} else {
		record(rc.clusterName, "ping", ERR)
	}
	return err
}
