//go:build ignore

// Lesson 6: Slices and maps.
//
// Run it with:
//
//	go run 06_slices_maps.go
package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	fmt.Println("--- Arrays ---")
	// An array has a fixed length that's part of its type: [3]int and [4]int
	// are different types. Arrays are values, so assigning one copies every
	// element. You'll rarely use arrays directly; slices are built on them.
	a := [3]int{1, 2, 3}
	b := a // a full copy
	b[0] = 100
	fmt.Println(a, b) // [1 2 3] [100 2 3]

	fmt.Println("--- Slices ---")
	// A slice is a resizable list: Go's version of ArrayList in Java,
	// List<T> in C#, or std::vector in C++. The literal looks like an
	// array's, but without a length in the brackets.
	s := []int{1, 2, 3}
	s = append(s, 4, 5) // append returns the updated slice; always assign it back
	fmt.Println(s, "length", len(s))
	fmt.Println("first", s[0], "last", s[len(s)-1])
	// Reading past the end panics ("index out of range"); you never get
	// garbage memory as you can in C.

	// make creates a slice of a given length, filled with zero values.
	zeros := make([]int, 3)
	fmt.Println(zeros) // [0 0 0]

	// Capacity is how many elements fit before append has to allocate a
	// bigger array and copy everything over. It roughly doubles each time.
	var grow []int
	var caps []int
	for i := range 10 {
		grow = append(grow, i)
		caps = append(caps, cap(grow))
	}
	fmt.Println("capacity after each append:", caps)

	// If you know the final size, set the capacity up front with make's third
	// argument to avoid the reallocations.
	squares := make([]int, 0, 5) // length 0, capacity 5
	for i := range 5 {
		squares = append(squares, i*i)
	}
	fmt.Println(squares, len(squares), cap(squares))

	fmt.Println("--- Slicing shares memory ---")
	// s[low:high] takes elements low through high-1. Either end can be left
	// out.
	nums := []int{10, 20, 30, 40, 50}
	middle := nums[1:4]
	fmt.Println(middle, nums[:2], nums[3:]) // [20 30 40] [10 20] [40 50]

	// middle is not a copy. It's a view into the same underlying array as
	// nums, so changing it changes nums.
	middle[0] = 999
	fmt.Println(nums) // [10 999 30 40 50]

	// The sneakier version: appending to a shorter slice can overwrite
	// elements of the longer one, because there's spare room after it.
	firstTwo := nums[:2]
	firstTwo = append(firstTwo, 777)
	fmt.Println(firstTwo, nums) // nums[2] is now 777

	// For an independent copy, use slices.Clone, or make plus copy.
	clone := slices.Clone(nums)
	clone[0] = -1
	// copy copies as many elements as fit and returns how many it copied.
	dst := make([]int, 2)
	copied := copy(dst, nums)
	fmt.Println(nums[0], clone[0], dst, copied) // 10 -1 [10 999] 2

	fmt.Println("--- Looping ---")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
	// fruit is a copy of each element. To change the slice, assign through
	// the index.
	for i := range fruits {
		fruits[i] = strings.ToUpper(fruits[i])
	}
	fmt.Println(fruits)

	fmt.Println("--- nil and empty slices ---")
	// The zero value of a slice is nil. A nil slice has length 0, and len,
	// range and append all work on it, so you rarely need to check for nil.
	var none []string
	empty := []string{}
	fmt.Println(none == nil, empty == nil, len(none), len(empty)) // true false 0 0
	none = append(none, "works")
	fmt.Println(none)

	fmt.Println("--- The slices package ---")
	scores := []int{42, 7, 19, 88, 7}
	slices.Sort(scores)
	fmt.Println(scores)                      // [7 7 19 42 88]
	fmt.Println(slices.Contains(scores, 19)) // true
	fmt.Println(slices.Index(scores, 42))    // 3
	fmt.Println(slices.Max(scores))          // 88

	fmt.Println("--- Maps ---")
	// A map is a hash table: HashMap in Java, Dictionary in C#,
	// std::unordered_map in C++.
	ages := map[string]int{
		"alice": 31,
		"bob":   25,
	}
	ages["carol"] = 40 // adds a key, or updates it if it exists
	fmt.Println(ages["alice"], len(ages))

	// Reading a missing key gives the zero value, not an error.
	fmt.Println(ages["nobody"]) // 0

	// To tell "missing" apart from "zero", ask for a second value.
	// By convention it's called ok.
	if age, ok := ages["dave"]; ok {
		fmt.Println("dave is", age)
	} else {
		fmt.Println("no dave")
	}

	delete(ages, "bob")
	fmt.Println(ages) // fmt prints maps sorted by key

	// Looping over a map visits the keys in random order, on purpose, so no
	// code can depend on the order. To go in order, sort the keys first.
	for _, name := range slices.Sorted(maps.Keys(ages)) {
		fmt.Println(name, ages[name])
	}

	// The zero value of a map is nil. Reading from a nil map works, but
	// writing to one panics, so create maps with a literal or with make.
	counts := make(map[string]int)
	text := "the cat and the hat and the bat"
	for _, word := range strings.Fields(text) { // Fields splits on whitespace
		counts[word]++ // a missing key starts at 0
	}
	fmt.Println("word counts:", counts)

	// Go has no set type. A map with bool values is the usual stand-in.
	seen := map[string]bool{}
	for _, word := range strings.Fields(text) {
		seen[word] = true
	}
	fmt.Println("contains hat?", seen["hat"], "contains dog?", seen["dog"])
}
