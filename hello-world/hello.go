package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
	"runtime"
	"github.com/stevejaxon/learning-go-lang/hello-world/stringutil"
	"strings"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func add(x int, y int) int {
	return x + y
}

// Note: When two or more consecutive named function parameters share a type, you can omit the type from all but the last
func sub(x, y int) int {
	return x-y
}

// Note: A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Note: Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// Note: A return statement without arguments returns the named return values. This is known as a "naked" return. (Naked return statements should be used only in short functions; they can harm readability in longer functions.)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Note: The var statement declares a list of variables; as in function argument lists, the type is last.
var foo, bar, foobar bool

// Note: A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var c, python, java = true, false, "no!"

// Note: variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	uni    rune       = '世'
)

// Note: Variables declared without an explicit initial value are given their zero value.
var i int     //0
var f float64 //0
var b bool    // false
var s string  // ""

// Note: The expression T(v) converts the value v to the type T. (Unlike in C, in Go assignment between items of different type requires an explicit conversion.)
var in int = 42
var fl float64 = float64(in)
var un uint = uint(fl)

// Note: Constants are declared like variables, but with the const keyword. Constants cannot be declared using the := syntax.
const World = "世界"

// Note: Go has only one looping construct, the for loop.
func forLoop() {
	sum := 0
	// The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
        // Note: Unlike other languages like C, Java, or Javascript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Note: The init and post statement are optional.
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// Note: C's 'while' is spelled for in Go.
	sum = 1
	for sum < 1025 {
		sum += sum
	}
	fmt.Println(sum)
}

// Note: Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
func ifStatement() {
	x := 8.0
	// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	if x > 0 {
		fmt.Println(math.Sqrt(x))
	}

	// Like for, the if statement can start with a short statement to execute before the condition. Variables declared by the statement are only in scope until the end of the if.
	if y:= math.Pow(3, 2); y < 10 {
		fmt.Println(y)	
	} 
}


func isNextDay(today time.Weekday, incrementBy time.Weekday) time.Weekday {
	return today + incrementBy
}

func switchStatement() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		// Note: A case body breaks automatically, unless it ends with a fallthrough statement.
		fallthrough
	case "fake":
		fmt.Println("Please Mind The Gap")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// Note: Switch cases can evaluate functions.
	// Note: Switch cases evaluate cases from top to bottom, stopping when a case succeeds. 
	switch time.Saturday {
	case isNextDay(today, 0):
		fmt.Println("Today.")
	case isNextDay(today, 1):
		fmt.Println("Tomorrow.")
	case isNextDay(today, 2):
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Note: Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Note: A defer statement defers the execution of a function until the surrounding function returns.
func deferStatement() {
	defer fmt.Println("world")

	fmt.Println("hello")

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Note: Go has pointers. A pointer holds the memory address of a value.
// Note: The type *T is a pointer to a T value. Its zero value is nil.
// Note: The & operator generates a pointer to its operand.
// Note: The * operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting".
// Note: Unlike C, Go has no pointer arithmetic.

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// Note: A struct is a collection of fields. (And a type declaration does what you'd expect.)
type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

	// Struct fields can be accessed through a struct pointer.

	// NOTE: To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p := &v
	p.Y = 1e9
	fmt.Println(v)

	// You can assign struct values using their name
	v1 := Vertex{Y: 162, X: 12 }
	fmt.Println(v1)
}


// Note: Arrays: the type [n]T is an array of n values of type T. The expression var a [10]int declares a variable a as an array of ten integers.
// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

func arrays() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = World
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Note: An array has a fixed size. A *slice*, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
	var slice []int = primes[1:4] // [1:4)
	var allSlice []int = primes[0:6] // [1:6) e.g. 1-5
	b := primes[0:1] // Can also be instantiated using :=
	// Will throw an out of bounds error: var allSlice []int = primes[0:7]
	fmt.Println(slice, allSlice, b)

	// Note: Slices are like references to arrays: A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(a, b)

	c[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	allSlice[5] = 17
	// Doesn't work because the range is not in the slice: b[5] = 17 : does compile though
	fmt.Println(primes)

	// Note: Slice literals: A slice literal is like an array literal without the length.
	// This is an array literal: [3]bool{true, true, false}
	// This creates the same array as above, then builds a slice that references it: []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// Note: Slice defaults : When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t) // len=6 cap=6 [2 3 5 7 11 13]

	t = t[1:4] // resulting in t = [3 5 7]
	fmt.Println(t)

	t = t[:2] // [0:2] resulting in t = [3 5]
	fmt.Println(t)

	t = t[1:] // [1:2] resulting in t = [5]
	fmt.Println(t)

	t = t[:] // [0:1] resulting in t = [5]
	fmt.Println(t)

	// Note: Slice length and capacity: A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	printSlice(t) // len=1 cap=4 [5]
	t = t[:4]
	printSlice(t) // len=4 cap=4 [5 7 11 13]
	t = t[2:4]
	printSlice(t) // len=2 cap=2 [11 13] -- can never get back to the original capacity

	// Note: Nil slices: The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
	var u []int
	v := []int{} // can't leave off the {} when using :=
	printSlice(u) // len=0 cap=0 []
	printSlice(v) // len=0 cap=0 []

	// If a Slice is of an explicitly created array then the underlying array is not affected by the slice being modified / truncated
	printSlice(allSlice)				// len=6 cap=6 [2 3 5 7 11 17] - same as primes array
	allSlice = allSlice[3:5]			// Truncates the slice's capacity and the 'view' of the array to just two elements
	printSlice(allSlice)				// len=2 cap=3 [7 11]
	allSlice = allSlice[:cap(allSlice)] // Can be re-sized back to len=3 cap=3 [7 11 17]
	printSlice(allSlice)
	fmt.Println(primes)					// [2 3 5 7 11 17] - has not been truncated

	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var sl []int
	printSlice(sl)
	if sl == nil {
		fmt.Println("nil!")
	}

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	sli := make([]int, 5)  // len=5 cap=5 [0 0 0 0 0]
	printSlice(sli)
	// To specify a capacity, pass a third argument to make:
	slic := make([]int, 12, 21)	// len=12 cap=21 [0 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(slic)

	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][1] = "O"
	board[2][0] = "X"
	board[1][1] = "O"
	board[2][1] = "X"

	// X O X
	// O O X
	// X X O
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// It is common to append new elements to a slice, and so Go provides a built-in append function.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
	// append works on nil slices.
	sl = append(sl, 0)
	printSlice(sl)						// len=1 cap=1 [0]
	// The slice grows as needed.
	sl = append(sl, 1, 2, 3, 4)
	printSlice(sl)						// len=5 cap=6 [0 1 2 3 4]

	allSlice = primes[:]
	printSlice(allSlice)	// len=6 cap=6 [2 3 5 7 11 17] - back to the complete primes array
	allSlice = append(allSlice, 19, 23, 29, 31, 37)	// len=11 cap=12 [2 3 5 7 11 17 19 23 29 31 37]
	printSlice(allSlice)
	fmt.Println(primes)		// [2 3 5 7 11 17] - the slice now points to a different array / the primes array has not been modified
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Exercise: Slices
// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
//	(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// (Use uint8(intValue) to convert between types.)
func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8((i^2 + j^2) ^ ((i * j) / 5) * 199)
		}
	}
	return ret
}

func rangeStatement() {
	// The range form of the for loop iterates over a slice or map.

	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// 2**0 = 1
	// 2**1 = 2
	// 2**2 = 4
	// 2**3 = 8
	// 2**4 = 16
	// 2**5 = 32
	// 2**6 = 64
	// 2**7 = 128
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	// If you only want the index, drop the ", value" entirely.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// You can skip the index or value by assigning to _.
	// 1
	// z2
	// 4
	// 8
	// 16
	// 32
	// 64
	// 128
	// 256
	// 512
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func mapStatement() {
	type Vertex struct {
		Lat, Long float64
	}

	// A map maps keys to values.

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	var m map[string]Vertex

	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}

	// Map literals are like struct literals, but the keys are required.
	var lit = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(lit) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var omit = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(omit) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// Insert or update an element in map :
	m["Bell Labs"] = Vertex{
		42.68433, -71.39967,
	}
	// Retrieve an element:
	fmt.Println(m["Bell Labs"]) // {42.68433 -71.39967}

	// Delete an element:
	delete(omit, "Google")
	fmt.Println(omit) // map[Bell Labs:{40.68433 -74.39967}]

	// Test that a key is present with a two-value assignment:
	v, ok := lit["Bell Labs"]
	fmt.Println(v, ok) // {40.68433 -74.39967} true
	v, ok = lit["Foo"]
	fmt.Println(v, ok) // {0 0} false
}

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsCounts := make(map[string]int)
	count := count(words)
	for _, word := range words {
		wordsCounts[word] = count[word]
	}
	return wordsCounts
}

func count(words []string) map[string]int {
	count := make(map[string]int)
	for _, word := range words {
		if _, ok := count[word]; ok {
			count[word] = count[word] + 1
		} else {
			count[word] = 1
		}
	}
	return count
}

// name (params) return type
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// The adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionStatement() {
	// Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot)) // 5
	fmt.Println(compute(math.Pow)) // 81

	// Go functions may be closures. A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// Seems to act a bit like an instance of a class with attributes - and the functions act upon only these attributes
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {

		// 0 0
		// 1 -2
		// 3 -6
		// 6 -12
		// 10 -20
		// 15 -30
		// 21 -42
		// 28 -56
		// 36 -72
		// 45 -90
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// Exercise: Fibonacci closure
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		temp := x
		x = y
		y = temp + y
		return x

		// or x, y = y, x + y
	}
}

func main() {
	// Using a user-created library
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	// Using the standard time library
	fmt.Println("The time is", time.Now())
	// Using the math library
	fmt.Println("My favorite number is", rand.Intn(43))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// Calling local functions
	fmt.Println(add(42, 13))
	fmt.Println(sub(100, 38))
	fmt.Println("returning multiple results from a function: ")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println("returning multiple, naked, results from a function: ")
	fmt.Println(split(17))
	// Variables
	foo = true;
	foobar = true;
	fmt.Println("printing the satate of the package level variables: ")
	fmt.Println(foo, bar, foobar)
	fmt.Println("printing the satate of the, initialised, package level variables: ")
	fmt.Println(c, python, java)
	fmt.Println("printing the satate of the, initialised, function level variables: ")
	var x, y, z = 100, 99.999999, "99.99998"
	fmt.Println(x, y, z)
	// Note: inside a function a shorthand for creating and initialising a variable exists; using ':='
	k := 3
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", uni, uni)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println("printing the converted / casted values: ")
	fmt.Println(in, fl, un)
	fmt.Println("Hello", World)
	fmt.Println("calling the for statements")
	forLoop()
	fmt.Println("calling the if statements")
	ifStatement()
	fmt.Println("calling the switch statements")
	switchStatement()
	fmt.Println("calling the defer statements")
	deferStatement()
	fmt.Println("calling the pointer statements")
	pointers()
	fmt.Println("calling the struct statements")
	structs()
	fmt.Println("calling the arrays statements")
	arrays()
	fmt.Println("calling the slices exercise")
	// requires <code>go get golang.org/x/tour/pic</code>
	pic.Show(Pic)
	fmt.Println("calling the range statements")
	rangeStatement()
	fmt.Println("calling the map statements")
	mapStatement()
	fmt.Println("calling the map exercise")
	// requires <code>go get golang.org/x/tour/wc</code>
	wc.Test(WordCount)
	fmt.Println("calling the function statements")
	functionStatement()
	fmt.Println("calling the function exercise")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
