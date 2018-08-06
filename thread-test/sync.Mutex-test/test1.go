package main

import "sync"

//已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	print(a, "证明f执行了\n")
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a, "main")
	l.Unlock()
}
