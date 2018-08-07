{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}

import (
	"golang.org/x/net/context"

	"github.com/stretchr/testify/mock"

	domain "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"
)

type Mock{{.Data.DbName}}Model struct {
	mock.Mock
}

func (self *Mock{{.Data.DbName}}Model) FindDemoById(ctx context.Context, id string) ([]*domain.Demo, error) {
	args := self.Called(ctx, ids)
	return args.Get(0).([]*domain.Demo), args.Error(1)
}

func (self *Mock{{.Data.DbName}}Model) HealthCheck() error {
	args := self.Called()
	return args.Error(0)
}

func (self *Mock{{.Data.DbName}}Model) Close() error {
	return nil
}

