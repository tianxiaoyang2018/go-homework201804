package config

import (
	"errors"

	"github.com/p1cn/tantan-backend-common/log"
)

var (
	ErrInvalidConfig = errors.New("invalid configuration")
	conf             Config
)

func GetCommonConfig() Config {
	return conf
}

// 公共配置
type Config struct {
	// 日志相关的配置
	Log log.Config
	// debug模块的配置
	Debug struct {
		// debug api监听的端口
		// interface:port
		Listen string
		// 是否是测试模式，测试模式会忽略一些校验逻辑
		Testing bool
		// 是否开启后门
		BackdoorEnabled bool
	}
	// api tracing
	Trace           TraceConfig
	GracefulTimeout Duration
}

type TraceConfig struct {
	// 关闭api tracing
	// 默认关闭
	Enable bool
	// 值可以是 tags 和 logs
	// tracing的时候是否收集tags和logs信息
	EnableParts []string // tags / logs
	// tracing收集途径
	Log    *log.Config
	Format string
}

func (self Config) IsValid() bool {
	// if len(self.Debug.Listen) == 0 {
	// 	return false
	// }
	return true
}

// 初始化公共组件
// 初始化日志
// 初始化trace
func InitCommon(cfg Config) error {

	if !cfg.IsValid() {
		return ErrInvalidConfig
	}
	conf = cfg

	err := log.Init(cfg.Log)
	if err != nil {
		return err
	}

	// 现在只有以log方式收集trace
	if cfg.Trace.Enable && cfg.Trace.Log != nil && len(cfg.Trace.Log.Output) > 0 {
		logger, err := log.NewLogger(*cfg.Trace.Log)
		if err != nil {
			log.Alert("initliaze trace log failed : err(%+v) , config(%+v)", err, cfg.Trace.Log)
			logger = log.GetLogger(log.DefaultLoggerName)
		}
		log.RegisterLogger(log.TraceLoggerName, logger)
	}

	return nil
}

type Log struct {
	Slog *log.Config
}
