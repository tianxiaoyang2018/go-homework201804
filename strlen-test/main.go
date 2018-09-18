package main

import (
	"fmt"
	"bytes"
	"strings"
	"unicode/utf8"
	"time"
)

func main() {
	str := "我是abc"
	fmt.Println("str = ",str)
	fmt.Println("按字节数打印，因一个汉字占4个字节，所以结果是9")
	fmt.Println("len(str) = ", len(str))
	fmt.Println("len([]byte(str)) = ", len([]byte(str)))

	fmt.Println("按int32打印，一个汉字就只占一个int32了，所以结果是5。rune是int32的别名")
	fmt.Println("len([]int32(str)) = ", len([]int32(str)))
	fmt.Println("len([]rune(str)) = ", len([]rune(str)))

	fmt.Println("其他方法")
	fmt.Println("bytes.Count([]byte(str),nil)-1 = ", bytes.Count([]byte(str),nil)-1)
	fmt.Println("strings.Count(str,\"\")-1=", strings.Count(str,"")-1)
	fmt.Println("strings.Count(str,\"\")-1=",utf8.RuneCountInString(str))



	str = "一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十"
	fmt.Println("\r\r 把str改为100个汉字，循环500W次，性能比较----")
	var strLen int
	var count = 5000000
	time1 := time.Now().UnixNano()
	for i:=0;i<count;i++ {
		strLen = len([]rune(str))
	}
	time2 := time.Now().UnixNano()
	fmt.Println("len([]rune(str)耗时", (time2 - time1)/1e6, "毫秒 ,字符串长度", strLen)


	time1 = time.Now().UnixNano()
	for i:=0;i<count;i++ {
		strLen = utf8.RuneCountInString(str)
	}
	time2 = time.Now().UnixNano()
	fmt.Println("utf8.RuneCountInString(str)耗时", (time2-time1)/1e6,"毫秒， 字符串长度",strLen)

	time1 = time.Now().UnixNano()
	for i:=0;i<count;i++ {
		strLen = bytes.Count([]byte(str),nil)-1
	}
	time2 = time.Now().UnixNano()
	fmt.Println("bytes.Count([]byte(str),nil)-1耗时", (time2-time1)/1e6,"毫秒， 字符串长度",strLen)

	time1 = time.Now().UnixNano()
	for i:=0;i<count;i++ {
		strLen = strings.Count(str,"")-1
	}
	time2 = time.Now().UnixNano()
	fmt.Println("strings.Count(str,\"\")-1耗时", (time2-time1)/1e6,"毫秒， 字符串长度",strLen)


	time1 = time.Now().UnixNano()
	for i:=0;i<count;i++ {
		strLen = len(str)
	}
	time2 = time.Now().UnixNano()
	fmt.Println("len(str)耗时", (time2-time1)/1e6,"毫秒， 字符串长度",strLen)

}
