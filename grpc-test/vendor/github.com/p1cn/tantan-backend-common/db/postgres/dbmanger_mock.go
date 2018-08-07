package postgres

import (
	"github.com/stretchr/testify/mock"
)

type MockDBManager struct {
	mock.Mock
}

func (self *MockDBManager) Get(op DbOperation) DBWrapper {
	args := self.Called(op)
	return args.Get(0).(DBWrapper)
}

func (self *MockDBManager) GetByNumber(op DbOperation, number int) (DBWrapper, int, error) {
	args := self.Called(op, number)
	return args.Get(0).(DBWrapper), args.Get(1).(int), args.Error(2)
}

func (self *MockDBManager) WalkShards(f func(db DBWrapper)) {
	self.Called(f)
}

func (self *MockDBManager) WithSchema(shardNumber int, queryTemplate string) (string, error) {
	args := self.Called(shardNumber, queryTemplate)
	return args.Get(0).(string), args.Error(1)
}

func (self *MockDBManager) GetByShardNumber(op DbOperation, shardNum int) (DBWrapper, error) {
	args := self.Called(op, shardNum)
	return args.Get(0).(DBWrapper), args.Error(1)
}

func (self *MockDBManager) GetPhysicalShardNumberByLogicalShardNumber(shardNum int) (int, error) {
	args := self.Called(shardNum)
	return args.Get(0).(int), args.Error(1)
}

func (self *MockDBManager) GetSchemaByShardNumber(shardNum int) string {
	args := self.Called(shardNum)
	return args.Get(0).(string)
}

func (self *MockDBManager) GetShardCount(number ...int) int {
	args := self.Called(number)
	return args.Get(0).(int)
}

func (self *MockDBManager) GetLogicalShardCount(number ...int) int {
	args := self.Called(number)
	return args.Get(0).(int)
}
