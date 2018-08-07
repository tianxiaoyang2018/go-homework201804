package config

// postgres 配置
type PostgreSql struct {
	// 监听的地址
	Address string
	// 监听的端口
	Port string
	// 用户名
	User string
	// 密码
	Password string
	// 数据库名字
	Database string
	// 数据库相关设置
	Settings string
	// 连接池大小
	PoolSize int
	// 从连接池获取连接超时时间
	PoolTimeout int
	// idle的超时时间
	IdleTimeout int
	// idle检查频率，秒
	IdleCheckFrequency int
	// 熔断器配置
	CircuitBreaker struct {
		// 是否关闭熔断器
		Disabled bool
		Type     string
		// 开路的间隔时间，单位秒
		// 开路后等待多少秒再进入半开路状态
		RetryDuration uint32
		// 半开路状态下，允许最大尝试次数
		MaxRequestsOnHalfOpenStatus uint32
		// 闭路状态下 清理错误count(TotalFailureThreshold)时间间隔，单位秒
		ClearCountInterval uint32
		// 闭路状态下 时间间隔内（ClearCountInterval）
		// 连续TotalFailureThreshold个错误的count则进入开路状态
		TotalFailureThreshold uint32
	}
}

// PostgresCluster
// Postgres集群配置
type PostgresCluster struct {
	// 逻辑分片
	Mod int
	// schema的前缀：例如 rel_8192_
	SchemaPrefix string
	// 分片配置：数组类型
	// 如果没有分片，则只需要填写一个元素
	Shards []struct {
		// 逻辑分片开始id
		FromLogicalShardMod int
		// 逻辑分片结束id
		ToLogicalShardMod int
		// master 配置
		Master PostgreSql
		// slaves的配置
		Slaves []PostgreSql
	}
}
