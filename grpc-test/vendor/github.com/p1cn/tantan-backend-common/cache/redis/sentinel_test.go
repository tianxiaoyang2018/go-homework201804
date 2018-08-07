package redis

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/p1cn/tantan-backend-common/config"
)

var sclient *sentinelClient

func TestNewSentinelClient(t *testing.T) {
	data, err := ioutil.ReadFile("sentinel.json")
	if err != nil {
		t.FailNow()
	}
	var config config.RedisConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		t.FailNow()
	}
	sclient, err = NewSentinelClient(config)
	if err != nil {
		t.FailNow()
	}
	assert.IsType(t, sentinelClient{}, *sclient)
}

func TestSentinelPing(t *testing.T) {
	err := sclient.Ping(context.Background())
	assert.NoError(t, err)
}

func TestSentinelSet(t *testing.T) {
	err := sclient.Set(context.Background(), "t", "tv")
	assert.NoError(t, err)

	err = sclient.Set(context.Background(), "t2", "tv2", time.Hour*48)
	assert.NoError(t, err)
}

func TestSentinelSets(t *testing.T) {
	kvs := make(map[string]interface{})
	kvs["t"] = "tv"
	kvs["t2"] = "tv2"
	err := hclient.Sets(context.Background(), kvs, time.Hour*48)
	assert.NoError(t, err)
}

func TestSentinelGet(t *testing.T) {
	v, err := sclient.Get(context.Background(), "t")
	assert.NoError(t, err)
	assert.Exactly(t, v, "tv")

	v, err = sclient.Get(context.Background(), "notexist")
	assert.Exactly(t, err, Nil)
}

func TestSentinelGets(t *testing.T) {
	vs, err := sclient.Gets(context.Background(), []string{"t", "t2"})
	assert.NoError(t, err)
	assert.Exactly(t, vs[0], "tv")
	assert.Exactly(t, vs[1], "tv2")
}

func TestSentinelExist(t *testing.T) {
	exists, err := sclient.Exist(context.Background(), "t")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = sclient.Exist(context.Background(), "notexist")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestSentinelExpire(t *testing.T) {
	err := sclient.Expire(context.Background(), "tv", -1)
	assert.NoError(t, err)

	exists, err := sclient.Exist(context.Background(), "tv")
	assert.NoError(t, err)
	assert.False(t, exists)
}
