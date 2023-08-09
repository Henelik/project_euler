package main

import (
	"fmt"
	"math"
	"strconv"
)

func is_prime(n int) bool {
	// fmt.Printf("Checking whether %v is prime\n", n)

	if n == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for d := 2; d <= limit; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func find_prime_family(n int, mask []bool) []int {
	prime_family := make([]int, 0, 10)
	n_string := strconv.Itoa(n)

	for i := 0; i < 10; i++ {
		// cull numbers where the replacement of digits changes the number of digits (0's at the beginning)
		if mask[0] && i == 0 {
			i++
		}

		s := []byte(n_string)

		for j, b := range mask {
			if b {
				s[j] = []byte(strconv.Itoa(i))[0]
			}
		}

		possible_prime, _ := strconv.Atoi(string(s))

		if is_prime(possible_prime) {
			prime_family = append(prime_family, possible_prime)
		}
	}

	return prime_family
}

func is_full(bit_mask []bool) bool {
	for _, b := range bit_mask {
		if !b {
			return false
		}
	}

	return true
}

func find_largest_prime_family(n int) []int {
	largest_family := []int{}
	bit_mask := make([]bool, len(strconv.Itoa(n)))

	for !is_full(bit_mask) {
		// increment the bit mask
		for b := len(bit_mask) - 1; b >= 0; b-- {
			if !bit_mask[b] {
				bit_mask[b] = true
				break
			}
			bit_mask[b] = !bit_mask[b]
		}

		result := find_prime_family(n, bit_mask)
		if len(result) > len(largest_family) {
			largest_family = result
		}
	}

	return largest_family
}

func main() {
	possible_solution := 2

	for {
		// only fully evaluate primes
		for !is_prime(possible_solution) {
			possible_solution++
		}

		result := find_largest_prime_family(possible_solution)
		if len(result) >= 8 {
			fmt.Printf("Found family for solution %v: %v\n", possible_solution, result)

			return
		}

		// fmt.Printf("Evaluated %v, family size %v\n", possible_solution, len(result))

		possible_solution++
	}
}
