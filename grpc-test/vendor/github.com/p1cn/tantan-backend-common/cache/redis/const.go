package redis

import (
	libredis "github.com/go-redis/redis"
)

const (
	// Nil reply Redis returns when key does not exist.
	Nil = libredis.Nil

	OK   = "ok"
	ERR  = "error"
	MISS = "miss"
)
