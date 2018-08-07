package postgres

import (
	"runtime"
	"strings"
)

// @TODO 为了运维和兼容，拉了这坨屎
// 后面将此逻辑移除掉，暴露接口给使用者自己计数
// 防止函数太多爆掉内存
var (
	functionMap = map[string]struct{}{}
)

func getFuncName(src string) string {
	lowerSrc := strings.ToLower(src)
	startIndex := 0
	for _, keyword := range []string{" from ", "select "} {
		if indexOfFrom := strings.Index(lowerSrc, keyword); indexOfFrom != -1 {
			startIndex = indexOfFrom + len(keyword) 
			dotPos := strings.Index(lowerSrc[startIndex:], ".")
			if dotPos != -1 {
				startIndex += dotPos + 1
			}
			break
		}
	}

	indexOfLeftParenthesis := strings.Index(lowerSrc[startIndex:], "(")
	if indexOfLeftParenthesis != -1 {
		SQL := "sql." + strings.TrimSpace(lowerSrc[startIndex:startIndex+indexOfLeftParenthesis])
		if functionNameReachMax(SQL) {
			return "error.max"
		}
		return SQL
	}

	pc, _, _, ok := runtime.Caller(5)
	if !ok {
		return ""
	}

	fnc := runtime.FuncForPC(pc)

	var name string
	list := strings.Split(fnc.Name(), ".")

	if len(list) > 1 {
		sname := list[len(list)-2]
		if len(sname) > 1 && sname[0] == '(' && sname[len(sname)-1] == ')' {
			sname = sname[1 : len(sname)-2]
			if len(sname) > 0 && sname[0] == '*' {
				sname = sname[1:]
			}
		}
		fname := list[len(list)-1]
		name = sname + "." + fname
	}

	goFunc := "go." + name

	if functionNameReachMax(goFunc) {
		return "error.max"
	}

	return goFunc
}

func functionNameReachMax(sql string) bool {
	_, ok := functionMap[sql]
	if ok {
		return false
	}
	if len(functionMap) < 100 {
		functionMap[sql] = struct{}{}
		return false
	}
	return true
}
