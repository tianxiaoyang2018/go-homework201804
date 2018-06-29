package main

import (
	"sync"
	"fmt"
)

var l sync.Mutex   // 只能加一次，不是可重入锁

func fun(n int) int {
	l.Lock()
	if n <= 2{
		l.Unlock()
		return n
	}else {
		sum := fun(n-1) + fun(n-2)
		l.Unlock()
		return sum
	}

}

func main() {
	fmt.Println(fun(3))
}
