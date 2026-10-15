//go:build ignore

// Lesson 9: Packages and modules.
//
// Run it with:
//
//	go run 09_packages.go
//
// A package is a folder of .go files compiled together. Every file in the
// folder starts with the same "package" line, and package main is the
// special one that becomes a program. A module is a set of packages with a
// go.mod file at the top. This folder's go.mod says "module learngo", so the
// package in the textutil folder is imported as "learngo/textutil".
// Open textutil/textutil.go to see the other side.
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"learngo/textutil"
)

func main() {
	fmt.Println("--- Using your own package ---")
	// You call a package's exported names as packagename.Name.
	fmt.Println(textutil.Reverse("stressed"))
	fmt.Println(textutil.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(textutil.WordCount("packages keep code organized"))

	// Names that start with a lowercase letter can't be used outside their
	// package. Uncomment the next line to see "undefined: textutil.normalize":
	// textutil.normalize("Hi")

	fmt.Println("--- The standard library ---")
	// Go comes with a large standard library: fmt, strings, os, time,
	// net/http and many more. You import a package by its path, and use the
	// last part of the path as its name: "unicode/utf8" is used as utf8.
	fmt.Println(strings.Repeat("ab", 3))
	fmt.Println(utf8.RuneCountInString("héllo"))

	// Documentation is written as comments above the code, and "go doc"
	// prints it. Try these in the terminal:
	//
	//	go doc learngo/textutil
	//	go doc learngo/textutil.IsPalindrome
	//	go doc strings.Fields
	//
	// The standard library's documentation is also at pkg.go.dev/std.

	// To use a third-party library, go get it. For example:
	//
	//	go get github.com/google/uuid
	//
	// That downloads the module and records it in go.mod (with its version)
	// and in go.sum (checksums of the download), then you can import
	// "github.com/google/uuid". "go mod tidy" adds anything missing from
	// go.mod and removes what's no longer used. Search for libraries at
	// pkg.go.dev.
}
