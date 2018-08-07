{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}


import (
	database "github.com/p1cn/tantan-backend-common/db/postgres"
	slog "github.com/p1cn/tantan-backend-common/log"
)

// storage interface
type {{.Data.PackageName}}Db interface {
	FindDemoById(string) ([]DemoPO, error)
}


func newDb(db database.DBManager) {{.Data.PackageName}}Db {
	return &{{.Data.PackageName}}DbAdapter{
		db: db,
	}
}

type {{.Data.PackageName}}DbAdapter struct {
	db database.DBManager
}

func (self *{{.Data.PackageName}}DbAdapter) FindDemoById(ctx context.Context, id string) ([]DemoPO, error) {
	var res DemosPO
	_, err := self.db.Get(database.DbRead).Query(&res, "sql", id)
	if err != nil {
		return nil, err
	}

	return res.C, nil
}


