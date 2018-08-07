{{.CopyRight}}
// package config
package config

import (
	"flag"
	"log"

	"github.com/p1cn/tantan-backend-common/version"
	common "github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/health"
)

{{if .Config.DB}}
const (
	{{range .Models}}
	DbName{{.DbName}} = "{{.DbName}}"
	{{end}}
)
{{end}}


var configPath = flag.String("config", "", "path to service config directory")
var globalConfig *Config

// Get - Get config
func Get() *Config {
	return globalConfig
}

func Init() *Config {
	// parse flags
	flag.Parse()

	// initiliaze version flag
	version.Init(common.ServiceName{{.ConstServiceName}})

	config := &Config{}

	configs := map[int]interface{}{
		common.CommonConfig:   &config.Common,
{{if .Config.DB}}		common.DatabaseConfig: &config.Database,{{end}}
{{if .Config.RPC}}		common.RpcConfig:      &config.RPC,{{end}}
{{if .Config.HTTP}}		common.HttpConfig:      &config.HTTP,{{end}}
{{if .Config.MQ}}		common.MqConfig:       &config.MQ,{{end}}
{{if .Config.DCL}}		common.DclConfig:      &config.DCL,{{end}} 
{{if .Config.Service}}		common.ServiceConfig:  &config.Service,{{end}}
{{if .Config.EventLog}}common.EventLogConfig:     &config.EventLog,{{end}}
    }

	err := common.ParseConfig(*configPath, configs)
	if err != nil {
		slog.Alert("%v", err)
		log.Fatal(err)
	}

	err = common.InitCommon(config.Common)
	if err != nil {
		slog.Alert("%v", err)
		log.Fatal(err)
	}

	if !config.IsValid() {
		log.Fatal(common.ErrInvalidConfig)
	}

	// initliazse health check
	health.Init(config.Common.Debug.Listen)

	globalConfig = config
	return config
}

// Config - Service config
type Config struct {
	Common   common.Config
{{if .Config.Service}}	Service  ServiceConfig{{end}}
{{if .Config.DB}}	Database common.Database{{end}}
{{if .Config.RPC}}	RPC      common.Rpc{{end}}
{{if .Config.HTTP}}	HTTP      common.Http{{end}}
{{if .Config.DCL}}	DCL      common.Dcl{{end}}
{{if .Config.MQ}}	MQ       common.MessageQueue{{end}}
{{if .Config.EventLog}}	EventLog       common.EventLog{{end}}

}

func (self Config) IsValid() bool {
	return true
}

type ServiceConfig struct {
}
