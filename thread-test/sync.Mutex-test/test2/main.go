package main

import (
	"fmt"

	"sync"

	"time"
)

var m *sync.Mutex

func main() {

	m = new(sync.Mutex)

	go lockPrint(1)

	lockPrint(2)

	time.Sleep(3 * time.Second)

	fmt.Printf("%s\n", "exit!")

}

func lockPrint(i int) {

	println(i, "lock start")

	m.Lock()

	println(i, "in lock")

	time.Sleep(3 * time.Second)

	m.Unlock()

	println(i, "unlock")

}
