package main

type Classic struct {
}

func (c Classic) GetTargetIndex(input []int, target int) int {
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
