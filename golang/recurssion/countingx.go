package main

// func main() {

// 	input := "axbxcxd"
// 	countingx(input)
// }

func countingx(input string) int {

	if len(input) == 0 {
		return 0
	}

	// if len(input) == 1 {
	// 	if string(input[0]) == "x" {
	// 		return 1
	// 	}

	// 	return 0
	// }
	if string(input[0]) == "x" {
		return 1 + countingx(input[1:])
	}

	return countingx(input[1:])
}
