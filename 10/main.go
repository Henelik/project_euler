package main

import "fmt"

func main() {
	primes := make([]int, 1, 2000000)
	primes[0] = 2
	sum := 2

	for i := 2; i < 2000000; i++ {
		is_prime := true
		for _, p := range primes {
			if i%p == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, i)
			sum += i
		}
	}

	fmt.Println(sum)
}
