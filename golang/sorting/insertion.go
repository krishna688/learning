package main

func insertion(input []int) []int {

	for i := 1; i < len(input); i++ {
		lowValue := input[i]

		for j := i - 1; j >= 0; j-- {
			if lowValue < input[j] {
				input[j], input[j+1] = lowValue, input[j]
			} else {
				break
			}
		}
	}

	return input
}
