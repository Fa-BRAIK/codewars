package main

import "fmt"

func main() {
	nums := []int{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5}

	fmt.Println(FindOdd(nums))
	fmt.Println(FindOddXor(nums))
}

func FindOdd(seq []int) (odd int) {
	nums := map [int]int{}

	for _, num := range seq {
		nums[num]++
	}

	for num, v := range nums {
		if v % 2 != 0 {
			odd = num; break
		}
	}

	return 
}

func FindOddXor(seq []int) (odd int) {
	for _, n := range seq {
		odd ^= n
	}

	return
}
