package main

import (
	"fmt"
	"strconv"
)

// is_prime determines if a number is prime, given a slice of the primes smaller than it
func is_prime(n int, primes []int) bool {
	if n == 0 {
		return false
	}

	for _, d := range primes {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func generate_primes(output chan<- int) {
	primes := []int{2}
	new_prime := 3

	output <- 2

	for {
		for !is_prime(new_prime, primes) {
			new_prime += 2
		}

		primes = append(primes, new_prime)
		new_prime += 2

		output <- primes[len(primes)-1]
	}
}

func are_concats_prime(a, b int, primes []int) bool {
	as := strconv.Itoa(a)
	bs := strconv.Itoa(b)
	concatA, _ := strconv.Atoi(as + bs)
	concatB, _ := strconv.Atoi(bs + as)

	return is_prime(concatA, primes) && is_prime(concatB, primes)
}

func main() {
	primesChan := make(chan int)
	primes := make([]int, 0, 10000)
	subset := make([]int, 5)

	go generate_primes(primesChan)

	for i := 0; i < 5; i++ {
		primes = append(primes, <-primesChan)
	}

	for {
		// set the last subset item to the last primes item to avoid duplicate searches
		subset[4] = primes[len(primes)-1]

		for a := 0; a < len(primes)-1; a++ {
			subset[0] = primes[a]

			if !are_concats_prime(subset[0], subset[4], primes) {
				continue
			}
			for b := a + 1; b < len(primes)-1; b++ {
				subset[1] = primes[b]

				switch {
				case !are_concats_prime(subset[1], subset[0], primes):
					continue
				case !are_concats_prime(subset[1], subset[4], primes):
					continue
				default:
				}

				for c := b + 1; c < len(primes)-1; c++ {
					subset[2] = primes[c]

					switch {
					case !are_concats_prime(subset[2], subset[0], primes):
						continue
					case !are_concats_prime(subset[2], subset[1], primes):
						continue
					case !are_concats_prime(subset[2], subset[4], primes):
						continue
					default:
					}

					for d := c + 1; d < len(primes)-1; d++ {
						subset[3] = primes[d]

						switch {
						case !are_concats_prime(subset[3], subset[0], primes):
							continue
						case !are_concats_prime(subset[3], subset[1], primes):
							continue
						case !are_concats_prime(subset[3], subset[2], primes):
							continue
						case !are_concats_prime(subset[3], subset[4], primes):
							continue
						default:
							sum := 0
							for _, n := range subset {
								sum += n
							}
							fmt.Printf("Found solution: %v, sum: %v\n", subset, sum)

							return
						}
					}
				}
			}
		}

		primes = append(primes, <-primesChan)
	}
}
