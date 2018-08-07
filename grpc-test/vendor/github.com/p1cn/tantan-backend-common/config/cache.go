package config

type Cache struct {
	Redis map[string]RedisConfig
}

type RedisAddr struct {
	Name string //不是一致性hash的时候可为空
	Addr string //host:port
}

type RedisConfig struct {
	//集群名称，用作监控时分类
	ClusterName string

	//cache类型
	CacheType string

	//master name for sentinel
	Name string

	// Map of name => host:port addresses of redis
	Addrs []RedisAddr

	// Hash algorithm, by default crc16(for hash")
	HashAlg string

	// Frequency of PING commands sent to check shards availability.
	// Shard is considered down after 3 subsequent failed checks.
	HeartbeatFrequency Duration

	DB       int
	Password string

	MaxRetries      int
	MinRetryBackoff Duration
	MaxRetryBackoff Duration

	DialTimeout  Duration
	ReadTimeout  Duration
	WriteTimeout Duration

	PoolSize           int
	PoolTimeout        Duration
	IdleTimeout        Duration
	IdleCheckFrequency Duration
}
