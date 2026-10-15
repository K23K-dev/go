//go:build ignore

// Lesson 8: Interfaces, plus a first look at generics.
//
// Run it with:
//
//	go run 08_interfaces.go
package main

import (
	"cmp"
	"errors"
	"fmt"
	"math"
)

// An interface is a list of methods. Any type that has those methods
// satisfies the interface automatically. There's no "implements" keyword:
// Circle and Rect below never mention Shape.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64      { return r.Width * r.Height }
func (r Rect) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// fmt.Stringer is a one-method interface from the fmt package:
//
//	type Stringer interface {
//		String() string
//	}
//
// When fmt prints a value that has a String method, it uses it, like
// toString() in Java or ToString() in C#. Rect has one; Circle doesn't.
func (r Rect) String() string {
	return fmt.Sprintf("%gx%g rectangle", r.Width, r.Height)
}

// totalArea accepts any mix of shapes.
func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// Optional: this line makes the compiler confirm that Circle satisfies
// Shape. If Circle lost a method, the program would stop compiling here.
var _ Shape = Circle{}

// any is an interface with no methods, so every type satisfies it, like
// Object in Java or object in C#. A type switch finds out what's inside.
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("the int %d", x)
	case string:
		return fmt.Sprintf("the string %q, %d bytes long", x, len(x))
	case Shape:
		return fmt.Sprintf("a shape with area %.2f", x.Area())
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("something else, of type %T", x)
	}
}

// Any type with an Error() string method is an error, so you can make your
// own error types that carry extra information.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Reason
}

func validateUsername(name string) error {
	if len(name) < 3 {
		return &ValidationError{Field: "username", Reason: "too short"}
	}
	// Return a plain nil for "no error". A classic trap: returning a
	// *ValidationError variable that happens to be nil gives the caller a
	// non-nil error, because the interface still records the pointer's type.
	return nil
}

// Generics: type parameters go in square brackets, each with a constraint
// saying what the type must support. The constraint "any" allows every type.
func Filter[T any](items []T, keep func(T) bool) []T {
	var result []T
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

// cmp.Ordered allows the types that support < and >: integers, floats and
// strings. Constraints are interfaces too. Another built-in constraint,
// comparable, allows the types that support ==.
func Largest[T cmp.Ordered](items []T) T {
	largest := items[0]
	for _, item := range items[1:] {
		if item > largest {
			largest = item
		}
	}
	return largest
}

// Types can have type parameters too.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop returns false as its second value when the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T // the zero value of whatever T is
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func main() {
	fmt.Println("--- Interfaces ---")
	shapes := []Shape{Circle{Radius: 1}, Rect{Width: 2, Height: 3}}
	for _, s := range shapes {
		fmt.Printf("%v: area %.2f, perimeter %.2f\n", s, s.Area(), s.Perimeter())
	}
	fmt.Printf("total area: %.2f\n", totalArea(shapes))

	fmt.Println("--- Type assertions ---")
	// An interface value holds a concrete value. A type assertion gets the
	// concrete value back out. With the second result (ok), a wrong guess
	// gives false instead of a panic.
	var shape Shape = Rect{Width: 4, Height: 5}
	if r, ok := shape.(Rect); ok {
		fmt.Println("it's a Rect with width", r.Width)
	}
	if _, ok := shape.(Circle); !ok {
		fmt.Println("it's not a Circle")
	}

	fmt.Println("--- any and type switches ---")
	for _, v := range []any{42, "hello", Circle{Radius: 2}, 3.5, nil} {
		fmt.Println(describe(v))
	}

	fmt.Println("--- Your own error types ---")
	err := validateUsername("al")
	// errors.As checks whether err, or any error wrapped inside it, is a
	// *ValidationError. If so, it stores it in ve so you can read its fields.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("problem with", ve.Field+":", ve.Reason)
	}
	fmt.Println("a valid name gives:", validateUsername("alice"))

	fmt.Println("--- Generics ---")
	// The type argument is usually inferred from the arguments.
	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4 6]
	long := Filter([]string{"go", "java", "c", "csharp"}, func(s string) bool { return len(s) > 2 })
	fmt.Println(long) // [java csharp]

	fmt.Println(Largest([]int{3, 9, 2}), Largest([]string{"pear", "apple", "zebra"}))

	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")
	top, _ := stack.Pop()
	fmt.Println("popped", top)

	// Like C# and unlike Java, Go doesn't erase the type at run time, so a
	// Stack[int] holds plain ints with no boxing. Unlike C++ templates, a
	// generic function is type-checked once against its constraint, not
	// separately for each type it's used with.
}
