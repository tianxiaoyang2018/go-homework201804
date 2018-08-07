package log

import (
	"github.com/stretchr/testify/mock"
)

type mockSyslogWriter struct {
	mock.Mock
}

func (self *mockSyslogWriter) Close() error {
	args := self.Called()
	return args.Error(0)
}

func (self *mockSyslogWriter) Flush() error {
	args := self.Called()
	return args.Error(0)
}

func (self *mockSyslogWriter) Alert(m string) error {
	args := self.Called(m)
	return args.Error(0)
}

func (self *mockSyslogWriter) Crit(m string) error {
	args := self.Called(m)
	return args.Error(0)
}

func (self *mockSyslogWriter) Debug(m string) error {
	args := self.Called(m)
	return args.Error(0)
}

func (self *mockSyslogWriter) Emerg(m string) error {
	args := self.Called(m)
	return args.Error(0)
}
func (self *mockSyslogWriter) Err(m string) error {
	args := self.Called(m)
	return args.Error(0)
}
func (self *mockSyslogWriter) Info(m string) error {
	args := self.Called(m)
	return args.Error(0)
}
func (self *mockSyslogWriter) Notice(m string) error {
	args := self.Called(m)
	return args.Error(0)
}
func (self *mockSyslogWriter) Warning(m string) error {
	args := self.Called(m)
	return args.Error(0)
}
