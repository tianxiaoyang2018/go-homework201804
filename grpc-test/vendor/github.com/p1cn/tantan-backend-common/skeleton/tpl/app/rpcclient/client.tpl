package rpcclient

import (
	common "github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/service-client"

    {{if .Config.Dependencies}}
    {{range .Deps}}
	"github.com/p1cn/tantan-domain-schema/golang/{{.LwConstServiceName}}"
	{{end}}
	{{end}}
)

{{if .Config.Dependencies}}
{{range .Deps}}
func New{{.ConstServiceName}}Client(name string, peer common.PeerService) ({{.LwConstServiceName}}.{{.ConstServiceName}}ServiceClient, error) {
	cfg := client.Config{
		ServiceName: name,
		Peer:        peer,
	}
	conn, err := client.NewConnection(cfg)
	if err != nil {
		return nil, err
	}

	{{.LwConstServiceName}}Client := {{.LwConstServiceName}}.New{{.ConstServiceName}}ServiceClient(conn)
	return {{.LwConstServiceName}}Client, nil
}
{{end}}
{{end}}
