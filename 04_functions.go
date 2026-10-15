//go:build ignore

// Lesson 4: Functions.
//
// Run it with:
//
//	go run 04_functions.go
package main

import (
	"fmt"
	"strings"
)

// Parameter types come after the names, and parameters of the same type
// can share it: (a, b int).
func add(a, b int) int {
	return a + b
}

// A function can return several values. This replaces out parameters in
// C#, small result classes in Java, and std::pair or std::tuple in C++.
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// Results can have names. They start at their zero values, and a bare
// "return" returns whatever they hold. Fine for short functions; in long
// ones, explicit return values are easier to follow.
// (This one assumes nums isn't empty.)
func minMax(nums []int) (lo, hi int) {
	lo, hi = nums[0], nums[0]
	for _, n := range nums[1:] {
		lo = min(lo, n) // min and max are built into the language
		hi = max(hi, n)
	}
	return
}

// A variadic function takes any number of arguments of one type.
// Inside the function, nums is a slice ([]int).
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Functions are values, so they can be passed around. The parameter type
// func(int) int means "a function that takes an int and returns an int".
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

// A function can also return a function. The returned function is a
// closure: it keeps using count after makeCounter has returned, and each
// call to makeCounter creates a separate count.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("--- Basics ---")
	fmt.Println(add(2, 3))
	fmt.Println("5! =", factorial(5))

	fmt.Println("--- Multiple return values ---")
	q, r := divmod(17, 5)
	fmt.Println("17 / 5 =", q, "remainder", r)

	// The blank identifier _ throws away a value you don't need. (Assigning
	// it to a variable you never use would be a compile error.)
	_, r = divmod(10, 3)
	fmt.Println("10 % 3 =", r)

	lo, hi := minMax([]int{4, -2, 9, 7})
	fmt.Println("min", lo, "max", hi)

	fmt.Println("--- Variadic functions ---")
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))
	nums := []int{10, 20, 30}
	fmt.Println(sum(nums...)) // ... passes a slice as the separate arguments

	fmt.Println("--- Functions as values ---")
	// An anonymous function stored in a variable, then one written inline.
	double := func(x int) int { return x * 2 }
	fmt.Println(applyTwice(double, 5))                           // 20
	fmt.Println(applyTwice(func(x int) int { return x + 3 }, 5)) // 11

	// Functions from other packages are values too.
	upper := strings.ToUpper
	fmt.Println(upper("quiet"))

	fmt.Println("--- Closures ---")
	next := makeCounter()
	fmt.Println(next(), next(), next()) // 1 2 3
	other := makeCounter()
	fmt.Println(other()) // 1: a separate count

	fmt.Println("--- defer ---")
	deferDemo()

	// Go has no function overloading and no default parameter values.
	// Use different names instead (like strings.Split and strings.SplitN),
	// or pass a struct holding the options.
}

// defer schedules a call to run when the surrounding function returns,
// however it returns. It's Go's version of finally in Java, using in C#,
// or a destructor in C++, and it's mostly used for cleanup, right after
// acquiring something:
//
//	f, err := os.Open("notes.txt")
//	if err != nil {
//		return err
//	}
//	defer f.Close()
func deferDemo() {
	fmt.Println("start")

	// The arguments are evaluated immediately; only the call waits.
	x := 1
	defer fmt.Println("deferred earlier: x was", x)
	x = 2

	// Deferred calls run in reverse order: the last one deferred runs first.
	for i := range 3 {
		defer fmt.Println("deferred in loop:", i)
	}

	fmt.Println("end of the function body, x is", x)
}
