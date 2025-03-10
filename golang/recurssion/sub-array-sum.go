package main

//	func main() {
//		input := []int{1, 2, 3, 4, 5, 6}
//		sum(input)
//		log.Println("final count", sum(input))
//	}
func sum(input []int) int {
	if len(input) == 1 {
		return input[0]
	}
	return input[0] + sum(input[1:])
}
