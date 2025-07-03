package main

import "fmt"

//	input := []interface{}{
//		1,
//		2,
//		[]interface{}{1, 2, 3, 4, []int{5, 6, 7}, []string{"kl", "kp"}},
//		6,
//		7,
//		9,
//		10,
//		[]int{1, 2, 4, 5, 6, 78, 8, 99, 23, 45, 67, 87, 89, 9},
//	}
func printValues(input []interface{}) {

	for _, val := range input {
		switch obj := val.(type) {
		case []int: // If it's a slice of ints, iterate over its elements
			for _, num := range obj {
				printValues([]interface{}{num})
			}
		case []string: // If it's a slice of ints, iterate over its elements
			for _, num := range obj {
				printValues([]interface{}{num})
			}
		case []interface{}:
			printValues(obj)
		default: // If it's a single value, print it directly
			// count++
			fmt.Println(val)
		}
	}

}
