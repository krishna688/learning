package main

import (
	"fmt"
	"log"
)

var (
	count int64
)

// generatePermutations generates all permutations without backtracking
func generatePermutations(arr []string) {
	var helper func([]string, []string)

	helper = func(remaining []string, current []string) {
		if len(remaining) == 0 {
			count++
			fmt.Println(current)
			return
		}

		for i, val := range remaining {
			next := append([]string{}, remaining[:i]...)
			next = append(next, remaining[i+1:]...)
			helper(next, append(current, val))
		}
	}

	helper(arr, []string{})
}

func main() {
	arr := []string{"a", "b", "c"}
	res := threePremutations(arr)
	for _, val := range res {
		fmt.Println(val)
	}

	log.Println(len(res))
}

func threePremutations(input []string) [][]string {
	result := [][]string{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			for k := 0; k < len(input); k++ {
				if i != j && i != k && k != j {
					res := []string{
						input[i], input[j], input[k],
					}
					result = append(result, res)
				}
			}
		}
	}

	return result

}
