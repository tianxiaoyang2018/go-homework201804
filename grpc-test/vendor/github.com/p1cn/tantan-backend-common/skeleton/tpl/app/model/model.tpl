{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}

import (
	"time"
	
	"golang.org/x/net/context"

	domain "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"

	database "github.com/p1cn/tantan-backend-common/db/postgres"
	{{if .DclCommiter}}
	"github.com/p1cn/tantan-backend-common/dcl"
	{{end}}
)

// Model interface providers some interfaces of database adpter (contains HealthCheck) 
type Model interface {
	HealthCheck() error
	Close() error
	FindDemoById(ctx context.Context, id string) ([]*domain.Demo, error)
}

// New model returns a new Model
// 
func NewModel(db database.DBManager{{if .DclCommiter}}, producer dcl_tools.Producer{{end}}) (Model, error) {

	db{{.Data.DbName}}, err := new{{.Data.DbName}}Db(db)
	if err != nil {
		return nil, err
	}

	{{if .DclCommiter}}
	return &dclCommiter{
		Model: &model{
			db{{.Data.DbName}}: db{{.Data.DbName}},
		}, 
		dclProducer: producer,
	}, nil
	{{else}}
	return &model{
				db{{.Data.DbName}}: db{{.Data.DbName}},
			}
	{{end}}
}



// model 
type model struct {
	db{{.Data.DbName}} {{.Data.PackageName}}Db
}

func (self *model) Close() error {
	return nil
}

func (self *model) HealthCheck() error {
	return nil
}

func (self *model) FindDemoById(ctx context.Context, id string) ([]*domain.Demo, error) {
	data, err := self.db{{.Data.DbName}}.FindDemoById(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDomainDemo(data), nil
}

func toDomainDemo(data []DemoPO) []*domain.Demo {  
}


