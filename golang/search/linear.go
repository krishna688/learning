package main

type linear struct{}

func (l linear) GetTargetIndex(input []int, target int) int {

	for i, val := range input {
		if val == target {
			return i
		}
	}

	return -1
}
