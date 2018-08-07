package config

type Config struct {
	ServiceName string `yaml:"AbbServiceName"`
	AppName     string `yaml:"AppName"`
	DebugPort   string `yaml:"DebugPort"`
	Cache       bool   `yaml:"Cache"`
	MQ          bool   `yaml:"MQ"`
	EventLog    bool   `yaml:"EventLog"`
	DclCommiter bool   `yaml:"DclCommiter"`

	Events []struct {
		Topic string `yaml:"Topic"`
	} `yaml:"Events"`
	Models []struct {
		DbName string `yaml:"DbName"`
	} `yaml:"Models"`
	RPC *struct {
		AbbName string `yaml:"AbbName"`
		Listen  string `yaml:"Listen"`
	} `yaml:"RPC"`
	HTTP *struct {
		Listen string `yaml:"Listen"`
	} `yaml:"HTTP"`
	Deps []struct {
		ServiceName    string `yaml:"ServiceName"`
		AbbServiceName string `yaml:"AbbServiceName"`
	} `yaml:"Deps"`
}
