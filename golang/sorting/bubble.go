package main

import "log"

func bubbleSort(input []int) []int {
	sorted_until_index := len(input) - 1
	sorted := false
	count := 0
	for !sorted {
		sorted = true
		for i := 0; i < sorted_until_index; i++ {
			count++
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				sorted = false
			}
		}
		sorted_until_index--
	}

	log.Println(count)
	return input
}
