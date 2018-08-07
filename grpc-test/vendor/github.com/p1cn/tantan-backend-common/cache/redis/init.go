package redis

import (
	"github.com/p1cn/tantan-backend-common/metrics"
)

type infoCollector struct {
	timer *metrics.Timer
}

var ic *infoCollector

func init() {
	ic = &infoCollector{
		timer: metrics.NewTimer(metrics.NameSpaceTantan, "cache_redis", "redis metrics", []string{"cluster_name", "op_name", "ret"}),
	}
}
