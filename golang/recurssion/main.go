package main

import (
	"log"
)

var count = 0

func main() {
	input := []interface{}{
		1,
		2,
		[]interface{}{1, 2, 3, 4, []int{5, 6, 7}, []string{"kl", "kp"}},
		6,
		7,
		9,
		10,
		[]int{1, 2, 4, 5, 6, 78, 8, 99, 23, 45, 67, 87, 89, 9},
	}

	printValues(input)
	log.Println("final count", count)
}
