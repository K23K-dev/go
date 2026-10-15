//go:build ignore

// Lesson 7: Structs, methods and pointers.
//
// Run it with:
//
//	go run 07_structs.go
package main

import (
	"errors"
	"fmt"
)

// A struct groups named fields. Go has no classes: you define a struct and
// attach methods to it. There's no inheritance either; see embedding below.
type Point struct {
	X, Y int
}

type Rectangle struct {
	Width, Height float64
}

// A method is a function with a receiver, written before the method name.
// The receiver plays the role of "this". This one is a value receiver, so
// the method gets a copy of the Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A pointer receiver (*Rectangle) gets a pointer to the original, so the
// method can change it. Use a pointer receiver when a method modifies the
// struct, or when the struct is big enough that copying it is wasteful.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor // short for (*r).Width: Go follows the pointer for you
	r.Height *= factor
}

// Fields starting with a lowercase letter are private to the package.
// Privacy in Go is per package, not per type: any code in this file could
// still read acct.balance directly. Other packages couldn't.
type Account struct {
	Owner   string
	balance int
}

// Go has no constructors. By convention, a function named NewSomething
// creates and returns a Something, often as a pointer.
func NewAccount(owner string) *Account {
	// Returning a pointer to a local variable is safe. The compiler sees
	// that the value outlives the function and allocates it on the heap,
	// and the garbage collector frees it once nothing points to it.
	return &Account{Owner: owner}
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

func (a *Account) Balance() int {
	return a.balance
}

// Embedding puts a type inside a struct without a field name. The embedded
// type's fields and methods are promoted, so you can use them directly on
// the outer struct. This is composition, not inheritance: a SavingsAccount
// isn't an Account, and you can't pass one where an *Account is expected.
type SavingsAccount struct {
	*Account
	Rate float64
}

func (s *SavingsAccount) AddInterest() {
	interest := int(float64(s.Balance()) * s.Rate)
	s.Deposit(interest) // Deposit is promoted from the embedded Account
}

func main() {
	fmt.Println("--- Creating structs ---")
	p1 := Point{X: 1, Y: 2} // with field names: the clearest form
	p2 := Point{3, 4}       // by position: allowed, but breaks if fields change
	var p3 Point            // the zero value: every field is zero
	fmt.Println(p1, p2, p3)
	fmt.Printf("%+v\n", p1) // %+v includes the field names

	p1.X = 10
	fmt.Println(p1.X)

	// Structs can be compared with ==, field by field.
	fmt.Println(p2 == Point{3, 4}) // true

	fmt.Println("--- Structs are values ---")
	// Assigning a struct copies it. In Java and C#, assigning an object
	// copies a reference instead. Go structs behave like C# structs or C++
	// objects.
	a := Point{1, 1}
	b := a
	b.X = 99
	fmt.Println(a, b) // {1 1} {99 1}

	fmt.Println("--- Pointers ---")
	// &x gives you a pointer to x, and *p reads or writes the value p points
	// to. There's no pointer arithmetic, and the garbage collector frees
	// memory, so there's no delete or free.
	n := 5
	ptr := &n
	*ptr = 6
	fmt.Println(n, *ptr) // 6 6

	// With a pointer to a struct, you can change the original.
	pa := &a
	pa.X = 42      // short for (*pa).X = 42
	fmt.Println(a) // {42 1}

	// A pointer's zero value is nil. Reading through a nil pointer panics.
	var missing *Point
	fmt.Println(missing == nil) // true

	fmt.Println("--- Methods ---")
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println("area:", r.Area())
	r.Scale(2) // Go passes &r automatically because Scale needs a pointer
	fmt.Println("after Scale(2):", r, "area:", r.Area())

	fmt.Println("--- A constructor function and a private field ---")
	acct := NewAccount("Alice")
	acct.Deposit(100)
	if err := acct.Withdraw(500); err != nil {
		fmt.Println("withdraw failed:", err)
	}
	fmt.Println(acct.Owner, "has", acct.Balance())

	fmt.Println("--- Embedding ---")
	savings := &SavingsAccount{Account: NewAccount("Bob"), Rate: 0.05}
	savings.Deposit(1000) // promoted from Account
	savings.AddInterest()
	fmt.Println(savings.Owner, "has", savings.Balance()) // 1050

	fmt.Println("--- Anonymous structs ---")
	// For one-off data you can skip declaring a named type.
	server := struct {
		Host string
		Port int
	}{"localhost", 8080}
	fmt.Printf("%+v\n", server)
}
