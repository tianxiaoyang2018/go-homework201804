package config

import (
	"time"
)

const (
	// service naming
	NamingTypeFile = "file"
	NamingTypeDns  = "dns"
)

// Database : 数据库配置
type Database struct {
	// Postgres 集群的配置(map[key]value)
	// key 为集群的业务名字，value为集群配置
	// e.g. : map["user"]PostgresCluster 则是user数据库集群的配置
	Postgres map[string]PostgresCluster
	// mysql *mysql.driver
}

// MessageQueue 消息队列配置
type MessageQueue struct {
	// Kafka 集群的配置(map[key]KafkaConfig)
	// key 为集群的业务名字
	Kafka map[string]KafkaConfig
	// RabbitMQ
}

// DCL 配置
type Dcl struct {
	// 引用的MQ类型
	MqType string
	// 引用的MQ集群的名字，对应MessageQueue配置map中的key
	MqName string
}

// 依赖服务的配置
type Dependencies struct {
	Services map[string]PeerService
}

// eventlog的配置
type EventLog struct {
	// 是否开启eventlog
	Enable bool
	// 消息队列的类型
	MqType string
	// 消息队列中map的key
	MqName string
}

type PeerServiceNaming struct {
	// 类型：
	// file : 从配置文件直接配置的方式来发现依赖的服务；Target则填写ip或者域名
	// consul : Target则填写consul地址
	Type   string
	Target string
}

// 依赖服务的配置
type PeerService struct {
	// 服务发现
	Naming PeerServiceNaming
	// 负载均衡策略, 默认roundrobin
	Balance *struct {
		Type string
	}
	// 服务标签
	Tags []string
	// 额外需要的中间件：common库支持的中间件名字列表
	Middleware []Middleware

	// 如果是grpc服务，grpc服务相关参数
	Grpc *GrpcClient
	//Thrift *ThriftClient
}

// 中间件配置
type Middleware struct {
	// 中间件名字
	Name string
	// 中间件的相关参数
	Meta map[string]interface{}
}

// RPC configuration
const (
	GrpcBalanceRR = "RoundRobin"
)

// grpc 配置
type Grpc struct {
	// 监听的接口和端口
	// interface:port
	Listen string
	// 中间件的名字; common库支持的
	Middleware []Middleware
}

type Rpc struct {
	Grpc *Grpc
}

type GrpcClient struct {
	Dial *GrpcDialOptions
	Call *GrpcCallOptions
}

type GrpcCallOptions struct {
	Timeout Duration
}

// grpc dial相关参数
type GrpcDialOptions struct {
	// 是否block
	WithBlock bool
	// block超时时间，WithBlock必须为true才生效
	WithTimeout Duration
	// insecure模式
	WithInsecure bool
}

type Http struct {
	Listen string
}

func DefaultGrpcClient() *GrpcClient {
	return &GrpcClient{
		Dial: &GrpcDialOptions{
			WithBlock:    false,
			WithTimeout:  Duration(3 * time.Second),
			WithInsecure: true,
		},
	}
}
