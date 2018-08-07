{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}


import (
	"github.com/stretchr/testify/mock"
)


type mock{{.Data.PackageName}}Db struct {
	mock.Mock
}

func (self *selfock{{.Data.PackageName}}Db) FindDemoById(id string) ([]demoPO, error) {
	args := self.Called(id)
	return args.Get(0).([]demoPO), args.Error(1)
}
