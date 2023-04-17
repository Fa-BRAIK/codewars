package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toWeirdCase("Weird string case"))
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
