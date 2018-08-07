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

var hclient *hashClient

func TestNewHashClient(t *testing.T) {
	data, err := ioutil.ReadFile("hash.json")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	var config config.RedisConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	hclient, err = NewHashClient(config)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	assert.IsType(t, hashClient{}, *hclient)
}

func TestHashPing(t *testing.T) {
	err := hclient.Ping(context.Background())
	assert.NoError(t, err)
}

func TestHashSet(t *testing.T) {
	err := hclient.Set(context.Background(), "t", "tv")
	assert.NoError(t, err)

	err = hclient.Set(context.Background(), "t2", "tv2", time.Hour*48)
	assert.NoError(t, err)
}

func TestHashSets(t *testing.T) {
	kvs := make(map[string]interface{})
	kvs["t"] = "tv"
	kvs["t2"] = "tv2"
	err := hclient.Sets(context.Background(), kvs, time.Hour*48)
	assert.NoError(t, err)
}

func TestHashGet(t *testing.T) {
	v, err := hclient.Get(context.Background(), "t")
	assert.NoError(t, err)
	assert.Exactly(t, v, "tv")

	v, err = hclient.Get(context.Background(), "notexist")
	assert.Exactly(t, err, Nil)
}

func TestHashGets(t *testing.T) {
	vs, err := hclient.Gets(context.Background(), []string{"t", "t2", "notexist"})
	assert.NoError(t, err)
	assert.Exactly(t, vs[0], "tv")
	assert.Exactly(t, vs[1], "tv2")
	assert.EqualValues(t, vs[2], Nil)
}

func TestHashExist(t *testing.T) {
	exists, err := hclient.Exist(context.Background(), "t")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = hclient.Exist(context.Background(), "notexist")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestHashExpire(t *testing.T) {
	err := hclient.Expire(context.Background(), "tv", -1)
	assert.NoError(t, err)

	exists, err := hclient.Exist(context.Background(), "tv")
	assert.NoError(t, err)
	assert.False(t, exists)
}
