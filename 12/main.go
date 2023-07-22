package main

import "fmt"

func count_divisors(n int) int {
	// start with 1 divisor: the number itself
	num_divisors := 1

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			num_divisors += 1
		}
	}

	return num_divisors
}

// completes in 6m34.17s
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
