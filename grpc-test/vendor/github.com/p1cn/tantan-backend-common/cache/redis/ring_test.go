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

var client *ringClient

func TestNewRingClient(t *testing.T) {
	data, err := ioutil.ReadFile("ring.json")
	if err != nil {
		t.FailNow()
	}
	var config config.RedisConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		t.FailNow()
	}
	client, err = NewRingClient(config)
	if err != nil {
		t.FailNow()
	}
	assert.IsType(t, ringClient{}, *client)
}

func TestPing(t *testing.T) {
	err := client.Ping(context.Background())
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	err := client.Set(context.Background(), "t", "tv")
	assert.NoError(t, err)

	err = client.Set(context.Background(), "t2", "tv2", time.Hour*48)
	assert.NoError(t, err)
}

func TestSets(t *testing.T) {
	kvs := make(map[string]interface{})
	kvs["t"] = "tv"
	kvs["t2"] = "tv2"
	err := hclient.Sets(context.Background(), kvs, time.Hour*48)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	v, err := client.Get(context.Background(), "t")
	assert.NoError(t, err)
	assert.Exactly(t, v, "tv")

	v, err = client.Get(context.Background(), "notexist")
	assert.Exactly(t, err, Nil)
}

func TestGets(t *testing.T) {
	vs, err := client.Gets(context.Background(), []string{"t", "t2", "notexist"})
	assert.NoError(t, err)
	assert.Exactly(t, vs[0], "tv")
	assert.Exactly(t, vs[1], "tv2")
	assert.EqualValues(t, vs[2], Nil)
}

func TestExist(t *testing.T) {
	exists, err := client.Exist(context.Background(), "t")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = client.Exist(context.Background(), "notexist")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestExpire(t *testing.T) {
	err := client.Expire(context.Background(), "tv", -1)
	assert.NoError(t, err)

	exists, err := client.Exist(context.Background(), "tv")
	assert.NoError(t, err)
	assert.False(t, exists)
}
