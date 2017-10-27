package main

import (
	"time"
	"fmt"
	"golang.org/x/tour/tree"
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

func fibonacciChannels(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func controllingChannelsExample() {
	c := make(chan int, 10)
	go fibonacciChannels(cap(c), c) // Also works if it is not run in a separate goroutine
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Println(i)
	}
}


func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectExample() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)

	// The default case in a select is run if no other case is ready.
	// Use a default case to try a send or receive without blocking:
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return // Breaks out of the for-loop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println(t.String())
	walkSubTree(t, ch)
	// Need to close the channel here
	close(ch)
}

func walkSubTree(t *tree.Tree, ch chan int, ) {
	if t == nil {
		return
	}
	walkSubTree(t.Left, ch)
	ch <- t.Value
	walkSubTree(t.Right, ch)
}

func WalkQuit(t *tree.Tree, ch, quit chan int) {
	walkSubTreeQuit(t, ch, quit)
	close(ch)
}

func walkSubTreeQuit(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkSubTreeQuit(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// Value successfully sent.
	case <-quit:
		return
	}
	walkSubTreeQuit(t.Right, ch, quit)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func printContentsOfChannel(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func channelExercise() {
	ch := make(chan int, 10)
	fmt.Println("Walk tree(1)")
	go Walk(tree.New(1), ch)
	printContentsOfChannel(ch)
	fmt.Println("Walk tree(2)")
	ch = make(chan int, 10)
	go Walk(tree.New(2), ch)
	printContentsOfChannel(ch)
	fmt.Println("Walk tree quit(3)")
	ch = make(chan int, 10)
	quit := make(chan int)
	defer close(quit)
	go WalkQuit(tree.New(3), ch, quit)
	printContentsOfChannel(ch)
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}

func main() {
	// GOROUTINE
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go <func> starts a new goroutine running
	fmt.Println("calling the goroutine statements")
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
	// NOTE: By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	fmt.Println("calling the channel statements")
	channelExamples()

	// BUFFERED CHANNELS
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
	// ch := make(chan int, 100)
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	fmt.Println("calling the buffered channel statements")
	bufferedChannelExamples()

	// CONTROLLING CHANNELS
	// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:
	// v, ok := <-ch
	// ok is false if there are no more values to receive and the channel is closed.
	fmt.Println("calling the controlling channels statements")
	controllingChannelsExample()

	// SELECT
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	fmt.Println("calling the select statements")
	selectExample()

	// CHANNEL EXERCISE
	// Requires "go get golang.org/x/tour/tree" to be run first
	fmt.Println("calling the channel exercise")
	channelExercise()
}
