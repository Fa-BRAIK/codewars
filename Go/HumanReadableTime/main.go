package main

import (
	"fmt"
)

func main() {
	fmt.Println(HumanReadableTime(0))
}

func HumanReadableTime(seconds int) string {
	h := int(seconds / 60 / 60)
	m := int((seconds - h * 60 * 60) / 60)
	s := seconds - (h * 60 * 60) - (m * 60)

	return fmt.Sprintf("%02d", h) + ":" + fmt.Sprintf("%02d", m) + ":" + fmt.Sprintf("%02d", s)
}
