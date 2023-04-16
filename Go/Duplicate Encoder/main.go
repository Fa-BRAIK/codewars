package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	word := os.Args[1]

	fmt.Println(DuplicateEncode(word))
}

func DuplicateEncode(word string) string {
	// transform word to lower case
	wordToLower := strings.ToLower(word)
	// Use a map to check if we already have a char
	// withint the word
	characters := make(map[rune]string)

	for _, char := range wordToLower {
		// If we have a character it means
		// we should replace it by ")"
		if _, ok := characters[char]; ok {
			characters[char] = ")"
			continue
		}

		// By default we replace by "("
		characters[char] = "("
	}

	// We build our string using our map
	str := ""
	for _, char := range wordToLower {
		str += characters[char]
	}

	return str
}
