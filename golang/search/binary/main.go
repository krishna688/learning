package main

import (
	"log"
)

func main() {

	input := []int{5, 23, 33, 54, 65, 72, 81, 83, 89, 92}

	for i, value := range input {
		log.Printf("index %d -> %d", i, value)
	}

	log.Println("\n output value index", binarySearch(input, 0))
}

func binarySearch(input []int, target int) int {
	left := 0
	right := len(input) - 1

	for left <= right {

		mid := left + (right-left)/2
		if input[mid] == target {
			return mid
		}

		if input[mid] < target {
			// target might be in the right
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
