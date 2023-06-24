package main

import "fmt"

func main() {
	for i := 20; i < 4294967295; i++ {
		valid := true
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println(i)
			break
		}
	}
}
