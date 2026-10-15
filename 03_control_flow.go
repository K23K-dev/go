//go:build ignore

// Lesson 3: Control flow.
//
// Run it with:
//
//	go run 03_control_flow.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--- if ---")
	// No parentheses around the condition, and the braces are required,
	// even for a single statement.
	temperature := 23
	if temperature > 30 {
		fmt.Println("hot")
	} else if temperature > 15 {
		fmt.Println("mild")
	} else {
		fmt.Println("cold")
	}

	// The condition must be a bool. Unlike in C and C++, "if count {" with
	// an int doesn't compile.

	// An if can start with a short statement, separated by a semicolon.
	// Variables declared there exist only inside the if and its else
	// branches. You'll see this constantly with errors (lesson 5).
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("parsed", n)
	}

	// Go has no ?: operator. Use an if instead.
	parity := "odd"
	if temperature%2 == 0 {
		parity = "even"
	}
	fmt.Println(temperature, "is", parity)

	fmt.Println("--- for is the only loop ---")
	// The classic three-part loop.
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// With only a condition, for is a while loop.
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("first power of 2 above 100:", n)

	// With nothing at all, it loops until a break or return.
	tries := 0
	for {
		tries++
		if tries == 3 {
			break
		}
	}
	fmt.Println("stopped after", tries, "tries")

	// range over a number counts from 0 up to, but not including, it.
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// range over a string gives each character's byte position and the
	// character itself.
	for i, ch := range "Go!" {
		fmt.Printf("%d=%c ", i, ch)
	}
	fmt.Println()

	// continue skips to the next iteration; break leaves the loop.
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// A label lets break and continue act on an outer loop. This stops
	// both loops at the first pair of numbers that adds up to 10.
	// ([]int{...} is a slice, Go's list type. Lesson 6 covers slices.)
	numbers := []int{3, 8, 4, 6}
outer:
	for i, a := range numbers {
		for _, b := range numbers[i+1:] {
			if a+b == 10 {
				fmt.Println("found", a, "+", b)
				break outer
			}
		}
	}

	fmt.Println("--- switch ---")
	// Cases don't fall through to the next one, so you don't write break.
	// A case can list several values.
	day := "sat"
	switch day {
	case "sat", "sun":
		fmt.Println(day, "is on the weekend")
	case "fri":
		fmt.Println(day, "is almost the weekend")
	default:
		fmt.Println(day, "is a weekday")
	}

	// A switch with no value tests each case as a condition. It's the usual
	// way to write a long if/else-if chain.
	score := 82
	switch {
	case score >= 90:
		fmt.Println("grade A")
	case score >= 80:
		fmt.Println("grade B")
	case score >= 70:
		fmt.Println("grade C")
	default:
		fmt.Println("grade F")
	}

	// To fall through on purpose, write fallthrough. It's rare in real code.
	switch level := 1; level {
	case 1:
		fmt.Println("level 1")
		fallthrough
	case 2:
		fmt.Println("level 2, reached by fallthrough")
	}

	fmt.Println("--- Putting it together: FizzBuzz ---")
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
