package template

import (
	"fmt"
	"strings"

	"github.com/p1cn/tantan-backend-common/skeleton/config"
)

type ConfigTplData struct {
	Cache        bool
	MQ           bool
	RPC          bool
	DB           bool
	DCL          bool
	HTTP         bool
	Dependencies bool
	EventLog     bool
	Service      bool
}

type EventTplData struct {
	Topic     string
	GroupName string
	GroupVar  string
	Handle    string
}

type ModelTplData struct {
	ModelClassName string
	PackageName    string
	DbName         string
	DclCommiter    bool
}

type ServiceTplData struct {
	ServiceClassName string
}

type RpcTplData struct {
	RpcInterface       string
	RpcServerClassName string
	RpcConstName       string
	RpcLwConstName     string
	Listen             string
}

type HttpTplData struct {
	Listen string
}

type DepTplData struct {
	ServiceName        string
	ConstServiceName   string
	LwConstServiceName string
}

type Model struct {
	AbbServiceName   string
	RepoName         string
	AppName          string
	ServiceName      string
	ConstServiceName string
	ServiceClassName string
	DebugPort        string
	CopyRight        string
	Config           ConfigTplData
	Events           []EventTplData
	Models           []ModelTplData
	DclCommiter      bool
	RPC              *RpcTplData
	HTTP             *HttpTplData
	Deps             []DepTplData
	Data             interface{}
}

func ModelFromConfig(cfg *config.Config) *Model {

	lwAbbServiceName := strings.ToLower(cfg.ServiceName[0:1]) + cfg.ServiceName[1:]
	serviceClassName := lwAbbServiceName + "Service"

	rpc := false
	if cfg.RPC != nil {
		rpc = true
	}

	db := false
	if len(cfg.Models) > 0 {
		db = true
	}
	dcl := false
	if len(cfg.Events) > 0 {
		dcl = true
	}
	dep := false
	if len(cfg.Deps) > 0 {
		dep = true
	}
	http := false
	if cfg.HTTP != nil {
		http = true
	}

	mq := false
	if len(cfg.Events) > 0 || cfg.DclCommiter {
		mq = true
	}

	events := []EventTplData{}
	for _, ev := range cfg.Events {
		name := strings.Replace(ev.Topic, "dcl.", "", -1)
		upName := strings.ToUpper(name[0:1]) + name[1:]
		events = append(events, EventTplData{
			Topic:     ev.Topic,
			GroupName: fmt.Sprintf("tantan-backend-%s-%s", lwAbbServiceName, ev.Topic),
			GroupVar:  name + "EventGroup",
			Handle:    fmt.Sprintf("Handle%sEvent", upName),
		})
	}

	models := []ModelTplData{}
	for _, md := range cfg.Models {
		lwName := strings.ToLower(md.DbName[0:1]) + md.DbName[1:]
		models = append(models, ModelTplData{
			ModelClassName: lwName + "Db",
			PackageName:    lwName,
			DbName:         md.DbName,
			DclCommiter:    true,
		})
	}

	deps := []DepTplData{}
	for _, dp := range cfg.Deps {
		lwName := strings.ToLower(dp.AbbServiceName[0:1]) + dp.AbbServiceName[1:]
		deps = append(deps, DepTplData{
			ServiceName:        fmt.Sprintf("tantan-backend-%s", lwName),
			ConstServiceName:   dp.AbbServiceName,
			LwConstServiceName: lwName,
		})
	}

	gbl := &Model{
		AbbServiceName:   cfg.ServiceName,
		AppName:          cfg.AppName,
		ServiceName:      cfg.ServiceName + "Service",
		RepoName:         fmt.Sprintf("tantan-backend-%s", lwAbbServiceName),
		ServiceClassName: serviceClassName,
		ConstServiceName: lwAbbServiceName,
		DebugPort:        cfg.DebugPort,
		CopyRight:        "// Copyright 2015 tantan Co., LTD . All Rights Reserved. ",
		Config: ConfigTplData{
			Cache:        cfg.Cache,
			MQ:           mq,
			RPC:          rpc,
			DB:           db,
			DCL:          dcl,
			HTTP:         http,
			Dependencies: dep,
			EventLog:     cfg.EventLog,
			Service:      true,
		},
		Events:      events,
		Models:      models,
		DclCommiter: cfg.DclCommiter,
		Deps:        deps,
	}

	if cfg.RPC != nil {
		lwName := strings.ToLower(cfg.RPC.AbbName[0:1]) + cfg.RPC.AbbName[1:]
		rpcData := &RpcTplData{
			RpcInterface:       cfg.RPC.AbbName + "Server",
			RpcServerClassName: lwName + "Server",
			RpcConstName:       cfg.RPC.AbbName,
			RpcLwConstName:     lwName,
			Listen:             cfg.RPC.Listen,
		}
		gbl.RPC = rpcData
	}

	if cfg.HTTP != nil {
		gbl.HTTP = &HttpTplData{
			Listen: cfg.HTTP.Listen,
		}
	}

	return gbl
}
