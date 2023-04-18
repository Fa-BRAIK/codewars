package main

import "fmt"

func main() {
	fmt.Println(ProductFib(5895))
	fmt.Println(ProductFibLinear(5895))
}

// O(n^2)
func ProductFib(prod uint64) (res [3]uint64) {
	fs := map[uint64]uint64{
		0: 0,
		1: 1,
	}
	var n uint64 = 2

	for {
		i := uint64(2)
		for i <= n + 1 {
			fs[i] = fs[i - 1] + fs[i - 2] 
			i++
		}

		f1 := fs[i - 2]
		f2 := fs[i - 1]

		mul := f1 * f2

		if mul == prod {
			res[0] = f1
			res[1] = f2
			res[2] = 1
			break
		}

		if mul > prod {
			res[0] = f1
			res[1] = f2
			res[2] = 0
			break
		}

		n++
	}

	return
}

func ProductFibLinear(prod uint64) [3]uint64 {
	f1 := uint64(0)
	f2 := uint64(1)

	for f1 * f1 < prod {
		f1, f2 = f2, f1 + f2
	}

	success := uint64(0)
	if prod == f1 * f2 {
		success = 1
	}

	return [3]uint64{f1, f2, success}
}
