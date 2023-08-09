package main

import (
	"fmt"
	"strconv"
)

func is_solution(n int) bool {
	n_string := strconv.Itoa(n)
	for i := 2; i <= 6; i++ {
		if !contains_same_characters(n_string, strconv.Itoa(n*i)) {
			return false
		}
	}
	return true
}

func contains_same_characters(a, b string) bool {
	a_map := create_character_map(a)
	b_map := create_character_map(b)

	for c, n := range a_map {
		if b_map[c] != n {
			return false
		}
	}

	return true
}

func create_character_map(s string) map[byte]int {
	result := map[byte]int{}
	for _, c := range []byte(s) {
		result[c]++
	}
	return result
}

func main() {
	for i := 1; true; i++ {
		if is_solution(i) {
			fmt.Println(i)
			return
		}
	}
}
