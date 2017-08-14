package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
        "math/cmplx"
	"runtime"
	"github.com/stevejaxon/learning-go-lang/hello-world/stringutil"
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
}
