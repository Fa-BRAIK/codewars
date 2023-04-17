package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(toWeirdCase("Weird string case"))
	fmt.Println(toWeirdCaseRunes("Weird string case"))
}

func toWeirdCase(str string) (wstr string) {
	toUpper := true
	for _, char := range str {
		if (char == ' ') {
			toUpper = true; wstr += " "
			continue
		}

		if toUpper {
			wstr += strings.ToUpper(string(char))
		} else {
			wstr += strings.ToLower(string(char))
		}
		toUpper = !toUpper
	}
	
	return
}

func toWeirdCaseRunes(str string) string {
	toUpper := true
	var result []rune
	for _, c := range str {
		if (c == 32) {
			result = append(result, c)
			toUpper = true
		} else if (toUpper) {
			result = append(result, unicode.ToUpper(c))
			toUpper = false
		} else {
			result = append(result, unicode.ToLower(c))
			toUpper = true
		}
	}
	
	return string(result)
}
