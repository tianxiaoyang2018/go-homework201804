package redis

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"

	"github.com/p1cn/tantan-backend-common/cache/redis/hashalg"
	"github.com/p1cn/tantan-backend-common/config"
)

type Hash func(data []byte) uint32

type hashClient struct {
	clusterName string

	clients []*redis.Client
	config  config.RedisConfig
	hash    Hash
}

var hashId int = 1

func NewHashClient(redisConfig config.RedisConfig) (*hashClient, error) {
	hc := &hashClient{
		config: redisConfig,
	}

	if redisConfig.ClusterName != "" {
		hc.clusterName = redisConfig.ClusterName
	} else {
		hc.clusterName = fmt.Sprintf("hash%d", hashId)
		hashId++
	}

	switch redisConfig.HashAlg {
	case "crc16", "":
		hc.hash = hashalg.Crc16sum
	default:
		return nil, errors.New("Invalid hash algorithm")
	}

	for _, redisAddr := range redisConfig.Addrs {
		option := redis.Options{
			Addr:         redisAddr.Addr,
			PoolTimeout:  time.Duration(redisConfig.PoolTimeout),
			PoolSize:     redisConfig.PoolSize,
			ReadTimeout:  time.Duration(redisConfig.ReadTimeout),
			WriteTimeout: time.Duration(redisConfig.WriteTimeout),
		}
		c := redis.NewClient(&option)
		hc.clients = append(hc.clients, c)
		if err := c.Ping().Err(); err != nil {
			return nil, err
		}
	}
	return hc, nil
}

func (hc *hashClient) Get(ctx context.Context, key string) (string, error) {
	record := ic.timer.Timer()
	c := hc.lookupShard(key)
	cmd := c.Get(key)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			record(hc.clusterName, "get", MISS)
		} else {
			record(hc.clusterName, "get", ERR)
		}
	} else {
		record(hc.clusterName, "get", OK)
	}
	return cmd.Result()
}

func (hc *hashClient) Exist(ctx context.Context, key string) (bool, error) {
	record := ic.timer.Timer()
	c := hc.lookupShard(key)
	cmd := c.Exists(key)
	if cmd.Err() != nil {
		record(hc.clusterName, "exist", OK)
		return false, cmd.Err()
	}
	record(hc.clusterName, "exist", OK)
	if cmd.Val() == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (hc *hashClient) Gets(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, errors.New("keys is empty")
	}
	record := ic.timer.Timer()
	clientMap := make([][]string, len(hc.clients))
	for _, key := range keys {
		pos := int(hc.hash([]byte(key))) % len(hc.clients)
		clientMap[pos] = append(clientMap[pos], key)
	}

	resMap := make(map[string]interface{})
	for pos, ks := range clientMap {
		if len(ks) == 0 {
			continue
		}
		results := hc.clients[pos].MGet(ks...)
		err := results.Err()
		if err != nil && err != redis.Nil {
			record(hc.clusterName, "gets", ERR)
			return nil, err
		}
		for i, v := range results.Val() {
			key := ks[i]
			resMap[key] = v
		}
	}

	res := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		value := resMap[key]
		if value == nil {
			record(hc.clusterName, "gets", MISS)
			value = Nil
		} else {
			record(hc.clusterName, "gets", OK)
		}
		res = append(res, value)
	}

	return res, nil
}

func (hc *hashClient) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	c := hc.lookupShard(key)
	var ex time.Duration
	if len(expiration) > 0 {
		ex = expiration[0]
	} else {
		ex = 0
	}
	record := ic.timer.Timer()
	status := c.Set(key, value, ex)
	if status.Err() != nil {
		record(hc.clusterName, "set", ERR)
		return status.Err()
	}
	record(hc.clusterName, "set", OK)
	return nil
}

func (hc *hashClient) Sets(ctx context.Context, kvs map[string]interface{}, expiration ...time.Duration) error {
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
	clientMap := make([]map[string]interface{}, len(hc.clients))
	for key, value := range kvs {
		pos := int(hc.hash([]byte(key))) % len(hc.clients)
		if clientMap[pos] == nil {
			orderMap := make(map[string]interface{})
			clientMap[pos] = orderMap
		}
		tempkvs := clientMap[pos]
		tempkvs[key] = value
	}

	for pos, tempkvs := range clientMap {
		if len(tempkvs) == 0 {
			continue
		}

		pipe := hc.clients[pos].Pipeline()
		pipelineCmds := make([]*redis.StatusCmd, 0, len(tempkvs))

		for key, value := range tempkvs {
			pipelineCmds = append(pipelineCmds, pipe.Set(key, value, ex))
		}
		_, err := pipe.Exec()
		if err != nil {
			record(hc.clusterName, "sets", ERR)
			return err
		}

		for _, pcmd := range pipelineCmds {
			_, err := pcmd.Result()
			if err != nil {
				record(hc.clusterName, "sets", ERR)
				return err
			}
		}
	}

	record(hc.clusterName, "sets", OK)
	return nil
}

func (hc *hashClient) Del(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return errors.New("keys is empty")
	}
	record := ic.timer.Timer()
	clientMap := make([][]string, len(hc.clients))
	for _, key := range keys {
		pos := int(hc.hash([]byte(key))) % len(hc.clients)
		clientMap[pos] = append(clientMap[pos], key)
	}

	for pos, ks := range clientMap {
		if len(ks) == 0 {
			continue
		}
		results := hc.clients[pos].Del(ks...)
		if results.Err() != nil {
			record(hc.clusterName, "del", ERR)
			return results.Err()
		}
	}
	record(hc.clusterName, "del", OK)
	return nil
}

func (hc *hashClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	record := ic.timer.Timer()
	c := hc.lookupShard(key)
	status := c.Expire(key, expiration)
	if status.Err() != nil {
		record(hc.clusterName, "expire", ERR)
		return status.Err()
	}
	record(hc.clusterName, "expire", OK)
	return nil
}

func (hc *hashClient) Ping(ctx context.Context) error {
	record := ic.timer.Timer()
	for _, c := range hc.clients {
		if err := c.Ping().Err(); err != nil {
			record(hc.clusterName, "ping", ERR)
			return err
		}
	}
	record(hc.clusterName, "ping", OK)
	return nil
}

func (hc *hashClient) lookupShard(key string) *redis.Client {
	if key == "" {
		return hc.randomShard()
	}
	return hc.clients[int(hc.hash([]byte(key)))%len(hc.clients)]
}

func (hc *hashClient) randomShard() *redis.Client {
	return hc.clients[rand.Int()%len(hc.clients)]
}
