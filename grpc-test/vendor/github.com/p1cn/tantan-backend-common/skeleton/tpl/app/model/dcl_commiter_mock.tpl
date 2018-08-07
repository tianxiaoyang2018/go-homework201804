{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}


import (
	"golang.org/x/net/context"

	"github.com/stretchr/testify/mock"

	"github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"
)

type mockDclCommiter struct {
	mock.Mock
}

func (self *mockDclCommiter) HealthCheck() error {
	args := self.Called()
	return args.Error(0)
}

func (self *mockDclCommiter) Close() error {
	args := self.Called()
	return args.Error(0)
}

// interface
func (self *mockDclCommiter) FindDemoById(ctx context.Context, id string) ([]*domain.Demo, error) {
	args := self.Called(ctx, ids)
	return args.Get(0).([]*domain.Demo), args.Error(1)
}
