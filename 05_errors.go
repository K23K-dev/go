//go:build ignore

// Lesson 5: Errors.
//
// Run it with:
//
//	go run 05_errors.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Go has no exceptions. A function that can fail returns an error as its
// last result, and the caller checks it. error is a built-in interface with
// a single method, Error() string. A nil error means success.

// A sentinel error is a predefined error value that callers can check for,
// much like catching a specific exception type. By convention the name
// starts with Err.
var ErrDivideByZero = errors.New("division by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

var ErrOutOfRange = errors.New("out of range")

// parseAge shows the usual pattern: call something that can fail, check the
// error right away, and return early if it failed.
func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		// %w wraps the original error inside a new one that adds context.
		// Callers can still find the original with errors.Is and errors.As.
		return 0, fmt.Errorf("invalid age %q: %w", s, err)
	}
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("age %d: %w", age, ErrOutOfRange)
	}
	return age, nil
}

// safeIndex shows recover turning a panic back into an ordinary error.
func safeIndex(nums []int, i int) (value int, err error) {
	defer func() {
		// recover only works inside a deferred function. It returns the
		// panic's value, or nil if there was no panic.
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	return nums[i], nil
}

func main() {
	fmt.Println("--- Returning and checking errors ---")
	result, err := divide(10, 4)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 / 4 =", result)
	}

	_, err = divide(1, 0)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Comparing with == works for a sentinel error that hasn't been wrapped.
	fmt.Println("is it ErrDivideByZero?", err == ErrDivideByZero)

	fmt.Println("--- Wrapping errors ---")
	for _, input := range []string{"42", "abc", "200"} {
		age, err := parseAge(input)
		if err != nil {
			fmt.Println("error:", err)

			// For wrapped errors, use errors.Is instead of ==. It checks the
			// error and everything wrapped inside it.
			if errors.Is(err, ErrOutOfRange) {
				fmt.Println("  errors.Is found ErrOutOfRange")
			}

			// errors.As finds an error of a specific type and gives you access
			// to its fields. strconv reports parse failures as *strconv.NumError.
			var numErr *strconv.NumError
			if errors.As(err, &numErr) {
				fmt.Printf("  errors.As found a *strconv.NumError for %q\n", numErr.Num)
			}
			continue
		}
		fmt.Println("age:", age)
	}

	fmt.Println("--- Ignoring an error ---")
	// You can discard an error with _, but only when failure really doesn't
	// matter. Here, a failed parse gives 0, which is acceptable.
	n, _ := strconv.Atoi("not a number")
	fmt.Println(n)

	fmt.Println("--- Making errors ---")
	fmt.Println(errors.New("something went wrong"))
	fmt.Println(fmt.Errorf("user %d not found", 7))

	fmt.Println("--- panic and recover ---")
	// panic is for bugs and for situations the program can't continue from,
	// such as reading past the end of a slice. It isn't try/catch: normal
	// failures are returned as errors. A panic unwinds the stack, running
	// deferred calls on the way, and crashes the program unless a deferred
	// function calls recover. You'll rarely write either.
	nums := []int{10, 20, 30}
	value, err := safeIndex(nums, 1)
	fmt.Println(value, err)
	value, err = safeIndex(nums, 5)
	fmt.Println(value, err)
}
