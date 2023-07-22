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

func worker(solution_chan chan<- int, start, stride, sum_start int) {
	triangle_sum := sum_start
	i := start

	for {
		for j := 0; j < stride; j++ {
			i += 1
			triangle_sum += i
		}

		if count_divisors(triangle_sum) > 500 {
			solution_chan <- triangle_sum
		}
	}
}

// completes in 0.058s
func main() {
	threads := 8
	triangle_sum := 0

	solution_chan := make(chan int)

	for t := 0; t < threads; t++ {
		triangle_sum += t
		go worker(solution_chan, t, threads, triangle_sum)
	}

	fmt.Printf("found solution: %v\n", <-solution_chan)
}
