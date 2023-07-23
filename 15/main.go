package main

import "fmt"

//This problem can be rephrased as
//"how many combinations of 0 and 1 of length 40 are there, with exactly 20 0s"

// count_permutations only counts the permutations starting with the given mode - HALF of all available
func count_permutations(num_current, num_other int) int {
	count := 0
	if num_other == 0 {
		return 1
	}
	for i := 1; i <= num_current; i++ {
		count += count_permutations(num_other, num_current-i)
	}
	return count
}

func main() {
	fmt.Println(count_permutations(20, 20) * 2)
}
