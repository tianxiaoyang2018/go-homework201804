{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}

import (
	"strconv"
	"time"

	pg "gopkg.in/pg.v3"

	database "github.com/p1cn/tantan-backend-common/db/postgres"
)

type DemoPO struct {
    Id  string  `pg:"id"`
    CreatedTime database.Time `pg:"created_time"`
}
