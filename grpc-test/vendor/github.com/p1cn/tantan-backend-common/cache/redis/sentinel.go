package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/p1cn/tantan-backend-common/config"
)

type sentinelClient struct {
	clusterName string

	client *redis.Client
	config config.RedisConfig
}

var sentinelId int = 1

func NewSentinelClient(redisConfig config.RedisConfig) (*sentinelClient, error) {
	sc := &sentinelClient{
		config: redisConfig,
	}

	if redisConfig.ClusterName != "" {
		sc.clusterName = redisConfig.ClusterName
	} else {
		sc.clusterName = fmt.Sprintf("sentinel%d", sentinelId)
		sentinelId++
	}

	sentinelAddrs := make([]string, 0, len(redisConfig.Addrs))

	for _, redisAddr := range redisConfig.Addrs {
		sentinelAddrs = append(sentinelAddrs, redisAddr.Addr)
	}

	option := redis.FailoverOptions{
		MasterName:    redisConfig.Name,
		Password:      redisConfig.Password,
		SentinelAddrs: sentinelAddrs,
		PoolTimeout:   time.Duration(redisConfig.PoolTimeout),
		PoolSize:      redisConfig.PoolSize,
		ReadTimeout:   time.Duration(redisConfig.ReadTimeout),
		WriteTimeout:  time.Duration(redisConfig.WriteTimeout),
        IdleTimeout:   time.Duration(redisConfig.IdleTimeout),
        MaxRetries:    redisConfig.MaxRetries,
        DialTimeout:   time.Duration(redisConfig.DialTimeout),
	}
	sc.client = redis.NewFailoverClient(&option)
	if err := sc.client.Ping().Err(); err != nil {
		return nil, err
	}
	return sc, nil

}

func (sc *sentinelClient) Get(ctx context.Context, key string) (string, error) {
	record := ic.timer.Timer()
	cmd := sc.client.Get(key)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			record(sc.clusterName, "get", MISS)
		} else {
			record(sc.clusterName, "get", ERR)
		}
	} else {
		record(sc.clusterName, "get", OK)
	}
	return cmd.Result()
}

func (sc *sentinelClient) Exist(ctx context.Context, key string) (bool, error) {
	record := ic.timer.Timer()
	cmd := sc.client.Exists(key)
	if cmd.Err() != nil {
		record(sc.clusterName, "exist", ERR)
		return false, cmd.Err()
	}
	record(sc.clusterName, "exist", OK)
	if cmd.Val() == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (sc *sentinelClient) Gets(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) <= 0 {
		return nil, errors.New("keys is empty")
	}
	record := ic.timer.Timer()
	res, err := sc.client.MGet(keys...).Result()
	if err != nil {
		record(sc.clusterName, "gets", ERR)
	} else {
		record(sc.clusterName, "gets", OK)
	}
	return res, err
}

func (sc *sentinelClient) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	var ex time.Duration
	if len(expiration) > 0 {
		ex = expiration[0]
	} else {
		ex = 0
	}
	record := ic.timer.Timer()
	status := sc.client.Set(key, value, ex)
	if status.Err() != nil {
		record(sc.clusterName, "set", ERR)
		return status.Err()
	}
	record(sc.clusterName, "set", OK)
	return nil
}

func (sc *sentinelClient) Sets(ctx context.Context, kvs map[string]interface{}, expiration ...time.Duration) error {
	if len(kvs) <= 0 {
		return errors.New("kvs is empty")
	}

	record := ic.timer.Timer()
	pairs := make([]interface{}, 0, len(kvs)*2)
	for key, value := range kvs {
		pairs = append(pairs, key)
		pairs = append(pairs, value)
	}

	status := sc.client.MSet(pairs...)
	if status.Err() != nil {
		record(sc.clusterName, "sets", ERR)
		return status.Err()
	}
	record(sc.clusterName, "sets", OK)
	return nil
}

func (sc *sentinelClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	record := ic.timer.Timer()
	status := sc.client.Expire(key, expiration)
	if status.Err() != nil {
		record(sc.clusterName, "expire", ERR)
		return status.Err()
	}
	record(sc.clusterName, "expire", OK)
	return nil
}

func (sc *sentinelClient) Del(ctx context.Context, keys []string) error {
	record := ic.timer.Timer()
	status := sc.client.Del(keys...)
	if status.Err() != nil {
		record(sc.clusterName, "del", ERR)
		return status.Err()
	}
	record(sc.clusterName, "del", OK)
	return nil
}

func (sc *sentinelClient) Ping(ctx context.Context) error {
	record := ic.timer.Timer()
	err := sc.client.Ping().Err()
	if err == nil {
		record(sc.clusterName, "ping", OK)
	} else {
		record(sc.clusterName, "ping", ERR)
	}
	return err
}
