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

var cClient *clusterClient

func TestNewClusterClient(t *testing.T) {
	data, err := ioutil.ReadFile("cluster.json")
	if err != nil {
		t.FailNow()
	}
	var config config.RedisConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		t.FailNow()
	}
	cClient, err = NewClusterClient(config)
	if err != nil {
		t.FailNow()
	}
	assert.IsType(t, clusterClient{}, *cClient)
}

func TestClusterPing(t *testing.T) {
	err := cClient.Ping(context.Background())
	assert.NoError(t, err)
}

func TestClusterSet(t *testing.T) {
	err := cClient.Set(context.Background(), "t", "tv")
	assert.NoError(t, err)

	err = cClient.Set(context.Background(), "t2", "tv2", time.Hour*48)
	assert.NoError(t, err)
}

func TestClusterSets(t *testing.T) {
	kvs := make(map[string]interface{})
	kvs["t"] = "tv"
	kvs["t2"] = "tv2"
	err := cClient.Sets(context.Background(), kvs, time.Hour*48)
	assert.NoError(t, err)
}

func TestClusterGet(t *testing.T) {
	v, err := cClient.Get(context.Background(), "t")
	assert.NoError(t, err)
	assert.Exactly(t, v, "tv")

	v, err = cClient.Get(context.Background(), "notexist")
	assert.Exactly(t, err, Nil)
}

func TestClusterGets(t *testing.T) {
	vs, err := cClient.Gets(context.Background(), []string{"t", "t2", "notexist"})
	assert.NoError(t, err)
	assert.Exactly(t, vs[0], "tv")
	assert.Exactly(t, vs[1], "tv2")
	assert.EqualValues(t, vs[2], Nil)
}

func TestClusterExist(t *testing.T) {
	exists, err := cClient.Exist(context.Background(), "t")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = cClient.Exist(context.Background(), "notexist")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestClusterExpire(t *testing.T) {
	err := cClient.Expire(context.Background(), "tv", -1)
	assert.NoError(t, err)

	exists, err := cClient.Exist(context.Background(), "tv")
	assert.NoError(t, err)
	assert.False(t, exists)
}
