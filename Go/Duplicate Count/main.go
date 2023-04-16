package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	word := os.Args[1]

	fmt.Println(DuplicateCount(word))
}

// Improved solution below
func DuplicateCount(s1 string) int {
	s1ToLower := strings.ToLower(s1)
	characters := make(map[rune]int)

	for _, char := range s1ToLower {
		if _, ok := characters[char]; ok {
			characters[char]++
			continue
		}

		// Should use Go zero value feature
		// Less code
		characters[char] = 1
	}

	fmt.Println(characters)

	count := 0
	for _, val := range characters {
		if val > 1 {
			count++
		}
	}

	return count
}

// declare a variable c that is going to be returned
// Make use of the zero value feature of Go 
// c is by default zero
func duplicate_count(s string) (c int) {
	characters := make(map [rune]int)

	for _, char := range strings.ToLower(s) {
		// We make use of the zero value feature of Go
		// so if we do ++ for the first time we are technically
		// incrementing a 0 to a 1
		if characters[char]++; characters[char] == 2 { c++ }
	}

	// telling the function to return c now
	return
}
