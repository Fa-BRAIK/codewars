package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	fmt.Println(DigPow(46288, 3))
}

func DigPow(n, p int) int {
	sum := 0
	for i, num := range strconv.Itoa(n) {
		sum += int(math.Pow(float64(num - '0'), float64(i + p)))
	}

	if sum % n == 0 {
		return sum / n
	}

	return -1
}