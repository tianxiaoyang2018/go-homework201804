{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}



import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	domain "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"

	database "github.com/p1cn/tantan-backend-common/db/postgres"
)


// unit test : FindDemoById interface
func TestFindDemoById(t *testing.T) {

	// ** mock database
	// new a demoDb
	mockDb := &mock{{.Data.DbName}}{}

	// response structure
	retDemos := []database.Demo{
		{
			ID: "1",
		},
	}

	// call FindDemoById("1") and return retDemos and error
	mockDb.On("FindDemoById", "1").Return(retDemos, nil)

	// ** new model with mock database
	m := model{
		db{{.Data.DbName}}: mockDb,
	}

	c, err := m.FindDemoById(context.Background(), "1")
	if err != nil {
		t.Error(err)
		return
	}

	mockDb.AssertExpectations(t)

	assert.Equal(t, c, []*domain.Demo{
		{
			Id: "1",
		},
	})
}
