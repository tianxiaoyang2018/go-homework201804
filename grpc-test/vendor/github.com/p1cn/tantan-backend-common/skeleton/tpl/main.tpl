{{.CopyRight}}
// package main
package main

import (
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/config"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/service"
	"github.com/p1cn/tantan-backend-common/runner"
)

func main() {
	cfg := config.Init()

	srv, err := service.New{{.ServiceName}}(cfg)
	if err != nil {
		slog.Alert("%v", err)
		return
	}

	runner.RunService(srv).Wait()
}
