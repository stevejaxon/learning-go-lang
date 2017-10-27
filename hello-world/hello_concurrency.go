package main

import (
	"time"
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
	"errors"
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

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	// We can also use defer to ensure the mutex will be unlocked as in the Value method.
	defer c.mux.Unlock()
	return c.v[key]
}

func mutexExample() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

//
func mutexExercise() {
	Crawl("http://golang.org/", 4, fetcher)
}

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()

	// We load it concurrently.
	body, urls, err := fetcher.Fetch(url)

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Done with %v\n", url)
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
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

	// SYNC.MUTEX
	// Go's standard library provides mutual exclusion with sync.Mutex and its two methods: Lock and Unlock
	fmt.Println("calling the mutex statements")
	mutexExample()

	// CHANNEL EXERCISE
	// Use Go's concurrency features to parallelise a web crawler.
	fmt.Println("calling the mutex exercise")
	mutexExercise()
}
