// Question 4: Maps and Sorting
// Write a program to sort a map by its keys. The map contains string keys and integer values.
package main

import (
	"sort"
)

func main() {

	input := map[string]int{
		"kp": 12,
		"b":  1,
		"a":  4,
		"5":  42,
	}

	output := []string{}
	for key := range input {
		output = append(output, key)
	}

	sort.Strings(output)
}
