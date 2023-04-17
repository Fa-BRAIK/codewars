package main

import "fmt"

func main() {
	recipe := map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	available := map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}

	fmt.Println(CackesRecursive(recipe, available))
}

func Cakes(recipe, available map[string]int) int {
	mina := 0

	for i, qty := range recipe {
		if qtya, ok := available[i]; ok {
			if qtya < qty {
				return 0
			}

			if amount := qtya / qty; amount < mina || mina == 0 {
				mina = amount
			}

			continue
		}

		return 0
	}

	return mina
}

func CackesRecursive(recipe, available map[string]int) int {
	for k, v := range recipe { 
		available[k] -= v

		if 0 > available[k] {
			return 0
		}
	}

	return 1 + CackesRecursive(recipe, available)
}
