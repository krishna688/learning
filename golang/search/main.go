package main

import (
	"log"
)

type Search interface {
	GetTargetIndex([]int, int) int
}

func main() {

	s, input, target := getObjNdInput("linear-ordered")

	log.Println("\n output value index", s.GetTargetIndex(input, target))
}

func getObjNdInput(alg string) (Search, []int, int) {
	log.Printf("running %s algorithm", alg)
	switch alg {
	case "rotated":
		return Rotated{}, []int{72, 81, 83, 89, 92, 5, 23, 33, 54, 65}, 23
	case "linear":
		return linear{}, []int{72, 81, 83, 89, 92, 5, 23, 33, 54, 65}, 23

	case "linear-ordered":
		return OrderedLinear{}, []int{5, 23, 33, 54, 65, 72, 81, 83, 89, 92}, 22
	default:
		return Classic{}, []int{5, 23, 33, 54, 65, 72, 81, 83, 89, 92}, 5

	}
}
