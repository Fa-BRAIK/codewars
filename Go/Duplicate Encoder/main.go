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
	wordToLower := strings.ToLower(word)
	characters := make(map [rune]string)
	
	for _, char := range wordToLower {
		if _, ok := characters[char]; ok {
			characters[char] = ")"
			continue
		} 

		characters[char] = "("
	}

	str := ""
	for _, char := range wordToLower {
		str += characters[char]
	}

	return str
}
