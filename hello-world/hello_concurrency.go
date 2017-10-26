package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// GOROUTINE
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go <func> starts a new goroutine running
	go say("world")
	say("hello")

	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	// The sync package provides useful primitives, although you won't need them much in Go as there are other primitives...

}
