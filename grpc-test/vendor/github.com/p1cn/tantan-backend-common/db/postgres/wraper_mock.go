package postgres

import (
	"github.com/stretchr/testify/mock"
	pg "gopkg.in/pg.v3"
)

type MockDBWrapper struct {
	mock.Mock
}

func (self *MockDBWrapper) Exec(q string, args ...interface{}) (pg.Result, error) {
	aa := []interface{}{
		q,
	}
	aa = append(aa, args...)
	rets := self.Called(aa...)
	return rets.Get(0).(pg.Result), rets.Error(1)
}

func (self *MockDBWrapper) ExecOne(q string, args ...interface{}) (pg.Result, error) {
	aa := []interface{}{
		q,
	}
	aa = append(aa, args...)
	rets := self.Called(aa...)
	return rets.Get(0).(pg.Result), rets.Error(1)
}

func (self *MockDBWrapper) Query(coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	aa := []interface{}{
		coll, q,
	}
	aa = append(aa, args...)
	rets := self.Called(aa...)
	return rets.Get(0).(pg.Result), rets.Error(1)
}

func (self *MockDBWrapper) QueryOne(record interface{}, q string, args ...interface{}) (pg.Result, error) {
	aa := []interface{}{
		record, q,
	}
	aa = append(aa, args...)
	rets := self.Called(aa...)
	return rets.Get(0).(pg.Result), rets.Error(1)
}

func (self *MockDBWrapper) RunInTransaction(fn func(tx *pg.Tx) error) error {
	rets := self.Called(fn)
	return rets.Error(0)
}
