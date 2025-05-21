package main

import (
	"fmt"
)

func main() {
	input := "believe it, you can achieve it! 世界"
	fmt.Println(ReserveAndPreserve(input))

}

func ReserveAndPreserve(input string) string {

	println(input)
	s := []rune(input)

	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer < rightPointer {

		if s[leftPointer] == ' ' {
			leftPointer += 1
		} else if s[rightPointer] == ' ' {
			rightPointer -= 1
		} else {
			s[leftPointer], s[rightPointer] = s[rightPointer], s[leftPointer]
			leftPointer += 1
			rightPointer -= 1
		}
	}

	return string(s)
}
