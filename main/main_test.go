package main_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/tianxiaoyang2018/go-test/main"
)

var _ = Describe("Main", func() {
	It("第一个测试用例", func() {
		main.Test()
	})
})
