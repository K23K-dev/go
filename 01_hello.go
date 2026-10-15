//go:build ignore

// Lesson 1: Hello, Go.
//
// Run it with:
//
//	go run 01_hello.go
//
// Every file in this folder starts with the line "//go:build ignore".
// Go normally compiles all the .go files in a folder together, as one
// program, so ten files with ten main functions would clash. That line
// leaves the file out of the folder's program. Naming the file on the
// command line (go run 01_hello.go) still runs it. Real projects don't need
// the line, because each program gets its own folder.
package main

// Imports list the packages this file uses: "fmt" for printing, "os" for
// command-line arguments and "runtime" for information about Go itself.
// An unused import is a compile error, not a warning. The VS Code Go
// extension fixes the import list for you when you save.
import (
	"fmt"
	"os"
	"runtime"
)

// main is where the program starts. Unlike in Java, C# or C++, it takes no
// parameters and returns nothing. Command-line arguments are in os.Args,
// and os.Exit(code) ends the program with an exit code.
func main() {
	// No semicolons: the compiler adds them at the ends of lines.
	fmt.Println("Hello, Go!")

	// Println puts spaces between its arguments and a newline at the end.
	fmt.Println("one", 2, 3.5, true)

	// Printf works like printf in C. %v ("value") prints anything in its
	// default format, and it's the verb you'll use most.
	name := "Alice"
	age := 30
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%v is %v years old\n", name, age)

	// A few more verbs:
	fmt.Printf("%q\n", name)             // "Alice": a quoted string
	fmt.Printf("%T, %T\n", name, age)    // string, int: the types
	fmt.Printf("%.2f\n", 3.14159)        // 3.14: two decimal places
	fmt.Printf("[%5d] [%-5d]\n", 42, 42) // [   42] [42   ]: width and alignment

	// Sprintf returns the string instead of printing it, like String.format
	// in Java or string.Format in C#.
	greeting := fmt.Sprintf("%s says hi", name)
	fmt.Println(greeting)

	// Println, Printf, Args and Version all start with a capital letter.
	// In Go that's what makes a name public (usable outside its package).
	// Lowercase names are private to their package. Lesson 9 covers this.
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS and CPU:", runtime.GOOS, runtime.GOARCH)

	// os.Args[0] is the program itself. "go run" compiles the file into a
	// temporary .exe and runs that, so this prints a path in a temp folder.
	// Try passing arguments: go run 01_hello.go apple banana
	fmt.Println("program:", os.Args[0])
	fmt.Println("arguments:", os.Args[1:])

	// To build a .exe you can keep, run: go build 01_hello.go
	// That creates 01_hello.exe in this folder. It's a single file that runs
	// on other Windows machines without Go or any runtime installed.

	// Go has one standard format for code, applied by a tool called gofmt.
	// VS Code runs it every time you save, so there's nothing to configure.
	// One rule the compiler itself enforces: an opening brace goes at the
	// end of the line, as in "func main() {". Putting the brace on the next
	// line, C# style, fails with "syntax error: unexpected semicolon or
	// newline before {", because the compiler adds a semicolon after
	// "func main()".
}
