package cache

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/p1cn/tantan-backend-common/config"
)

func TestFactory(t *testing.T) {
	data, err := ioutil.ReadFile("redis/ring.json")
	if err != nil {
		t.FailNow()
	}
	var config1 config.RedisConfig
	err = json.Unmarshal(data, &config1)
	if err != nil {
		t.FailNow()
	}
	client, err := Factory(config1.CacheType, config1)
	assert.NoError(t, err)
	assert.Implements(t, (*ICacheClient)(nil), client)

	data, err = ioutil.ReadFile("redis/sentinel.json")
	if err != nil {
		t.FailNow()
	}
	var config2 config.RedisConfig
	err = json.Unmarshal(data, &config2)
	if err != nil {
		t.FailNow()
	}
	client, err = Factory(config2.CacheType, config2)
	assert.NoError(t, err)
	assert.Implements(t, (*ICacheClient)(nil), client)

	data, err = ioutil.ReadFile("redis/cluster.json")
	if err != nil {
		t.FailNow()
	}
	var config3 config.RedisConfig
	err = json.Unmarshal(data, &config3)
	if err != nil {
		t.FailNow()
	}
	client, err = Factory(config3.CacheType, config3)
	assert.NoError(t, err)
	assert.Implements(t, (*ICacheClient)(nil), client)

	data, err = ioutil.ReadFile("redis/hash.json")
	if err != nil {
		t.FailNow()
	}
	var config4 config.RedisConfig
	err = json.Unmarshal(data, &config4)
	if err != nil {
		t.FailNow()
	}
	client, err = Factory(config4.CacheType, config4)
	assert.NoError(t, err)
	assert.Implements(t, (*ICacheClient)(nil), client)
}
