package main_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/tianxiaoyang2018/go-test/main"
	"fmt"
)

var _ = Describe("第一批", func() {
	It("第二个测试用例", func() {
		main.Test()
		fmt.Println("objk")
	})
})
