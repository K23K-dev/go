//go:build ignore

// Lesson 2: Variables, types and constants.
//
// Run it with:
//
//	go run 02_variables.go
package main

import "fmt"

// Constants are declared with const. A constant without a type, like this
// one, is exact until it's used, then takes whatever numeric type the
// context needs.
const Pi = 3.14159

// Go has no enum keyword. Instead you define a type and a block of
// constants, and iota numbers them 0, 1, 2, ... Weekday is a new type
// based on int, so a Weekday can't be mixed up with a plain int variable.
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1: the type and "= iota" carry over
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// iota also works inside expressions. Assigning to _ skips a value.
const (
	_  = iota             // 0, thrown away
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20
	GB                    // 1 << 30
)

func main() {
	fmt.Println("--- Declaring variables ---")

	// The full form is "var name type = value": the type comes after the
	// name.
	var count int = 10

	// With a value, the type can be left out and Go infers it, like var in
	// Java and C# or auto in C++.
	var ratio = 0.75 // float64

	// Inside functions, := declares and assigns in one step. It's the form
	// you'll see most.
	city := "Lisbon"

	// Several at once. Assigning to several variables at once also gives
	// you a swap without a temporary variable.
	x, y := 1, 2
	x, y = y, x
	fmt.Println(count, ratio, city, x, y)

	fmt.Println("--- Zero values ---")
	// A variable declared without a value gets its type's zero value. You
	// never get uninitialized memory, and these types can't be null.
	var i int
	var f float64
	var s string
	var b bool
	fmt.Printf("int=%v float64=%v string=%q bool=%v\n", i, f, s, b)

	fmt.Println("--- Basic types ---")
	// int is 64 bits on a 64-bit system. When you need a specific size,
	// there are int8, int16, int32 and int64, and unsigned uint8 to uint64.
	// float64 is the usual floating-point type; float32 also exists.
	var big int64 = 9_000_000_000 // underscores make long numbers readable
	var small uint8 = 255
	var price float64 = 19.99
	var done bool = true
	fmt.Println(big, small, price, done)

	// byte is another name for uint8. rune is another name for int32 and
	// holds one Unicode character. Character literals use single quotes.
	var letter byte = 'A'
	var heart rune = '♥'
	fmt.Println(letter, heart)                 // 65 9829: they're numbers
	fmt.Println(string(letter), string(heart)) // A ♥

	fmt.Println("--- Conversions are always explicit ---")
	// Go never converts between number types for you, not even from int to
	// float64. Uncomment the next line to see the compile error
	// "invalid operation: count + ratio (mismatched types int and float64)":
	// fmt.Println(count + ratio)
	total := float64(count) + ratio
	fmt.Println(total)

	// Converting a float to an int drops the fraction.
	pi := 3.99
	fmt.Println(int(pi), int(-pi)) // 3 -3

	// Dividing two integers drops the remainder, as in Java, C# and C++.
	fmt.Println(7/2, 7.0/2, float64(7)/2) // 3 3.5 3.5

	// Careful: string(65) gives "A" (the character with code 65), not "65".
	// For the digits, use fmt.Sprint or strconv.Itoa.
	fmt.Println(string(rune(65)), fmt.Sprint(65)) // A 65

	fmt.Println("--- Unused variables are compile errors ---")
	// A local variable that's never used is an error, not a warning.
	// Uncomment the next line to see "declared and not used: unused":
	// unused := 42
	fmt.Println("(nothing to print; see the comment in the code)")

	fmt.Println("--- Constants and enums ---")
	fmt.Println(Pi * 2)
	fmt.Println(Sunday, Monday, Saturday) // 0 1 6
	today := Wednesday
	fmt.Println("is today Saturday?", today == Saturday)
	fmt.Println(KB, MB, GB)

	fmt.Println("--- Strings ---")
	// A string is an immutable sequence of bytes, usually UTF-8 text.
	// len counts bytes, not characters.
	word := "héllo"
	fmt.Println(len(word))         // 6: é takes 2 bytes in UTF-8
	fmt.Println(len([]rune(word))) // 5 characters

	// Indexing gives you a byte. Looping with range gives you characters
	// (runes), along with the byte position where each one starts.
	fmt.Println(word[0], string(word[0])) // 104 h
	for index, char := range word {
		fmt.Printf("%d:%c ", index, char) // the position jumps from 1 to 3 after é
	}
	fmt.Println()

	// Backquotes make a raw string: no escape sequences, and it can span
	// several lines. Handy for Windows paths.
	path := `C:\Program Files\Go\bin`
	fmt.Println(path)

	// + joins strings. You can't change a string in place; you build a new one.
	fmt.Println("go" + "pher")
}
