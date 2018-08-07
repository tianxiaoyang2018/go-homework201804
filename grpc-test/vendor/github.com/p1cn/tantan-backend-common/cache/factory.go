package cache

import (
	"encoding/json"
	"io/ioutil"

	"github.com/p1cn/tantan-backend-common/cache/redis"
	"github.com/p1cn/tantan-backend-common/config"
)

//根据配置配置初始化cacheClient
func FactoryByConfig(configPath string) (ICacheClient, error) {
	//读取cache的配置文件并初始化
	var cacheClientConfig config.RedisConfig
	if err := parseObject(configPath, &cacheClientConfig); err != nil {
		return nil, err
	}

	//创建cacheclient
	cacheClient, err := Factory(cacheClientConfig.CacheType, cacheClientConfig)
	if err != nil {
		return nil, err
	}

	return cacheClient, nil
}

//解析配置文件
func parseObject(file string, obj interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

// @todo : 加监控 （必须包含：redis业务集群名字）
// cacheType: redis_sentinel, redis_ring
func Factory(cacheType string, cfg interface{}) (ICacheClient, error) {
	var client interface{}
	var err error
	switch cacheType {
	case "redis_sentinel":
		client, err = redis.NewSentinelClient(cfg.(config.RedisConfig))
		if err != nil {
			return nil, err
		}
	case "redis_ring":
		client, err = redis.NewRingClient(cfg.(config.RedisConfig))
		if err != nil {
			return nil, err
		}
	case "redis_cluster":
		client, err = redis.NewClusterClient(cfg.(config.RedisConfig))
		if err != nil {
			return nil, err
		}
	case "redis_hash":
		client, err = redis.NewHashClient(cfg.(config.RedisConfig))
		if err != nil {
			return nil, err
		}
	}
	return client.(ICacheClient), nil
}
