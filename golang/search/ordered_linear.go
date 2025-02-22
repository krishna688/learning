package main

type OrderedLinear struct {
}

func (o OrderedLinear) GetTargetIndex(input []int, target int) int {
	for i, val := range input {

		if val == target {
			return i
		}
		if val > target {
			break
		}
	}

	return -1
}
