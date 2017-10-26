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

// Note: Channels can be used as function parameters / arguments
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channelExamples() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// Split the work between two goroutines - slicing the array to achieve the portioning of the work
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func bufferedChannelExamples() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 - would cause the code to block here (fatal error: all goroutines are asleep - deadlock!)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	// GOROUTINE
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go <func> starts a new goroutine running
	go say("world")
	say("hello")

	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	// The sync package provides useful primitives, although you won't need them much in Go as there are other primitives...

	// CHANNELS
	// Channels are a typed conduit through which you can send and receive values with the channel operator, <-. (The data flows in the direction of the arrow.)
	// e.g.
	// ch <- v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and assign value to v.

	// Like maps and slices, channels must be created before use:
	// ch := make(chan int)
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	channelExamples()

	// BUFFERED CHANNELS
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
	// ch := make(chan int, 100)
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	bufferedChannelExamples()
}
