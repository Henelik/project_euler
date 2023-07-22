package main

import (
	"fmt"
	"math"
)

func count_divisors(n int) int {
	// start with 1 divisor: the number itself
	num_divisors := 1

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			num_divisors += 2
		}
	}

	return num_divisors
}

// completes in 0.147s
func main() {
	triangle_sum := 0

	for i := 0; true; i += 1 {
		triangle_sum += i
		divisors := count_divisors(triangle_sum)
		//fmt.Println(triangle_sum, " has ", divisors, " divisors; index: ", i)
		if divisors > 500 {
			fmt.Printf("found solution: %v\n", triangle_sum)
			break
		}
	}
}
