package main

// func main() {
// 	input := "krishna prasad"

//		log.Println(stringReversal(input))
//	}
func stringReversal(input string) string {
	if len(input) == 1 {
		return input
	}
	return stringReversal(input[1:]) + string(input[0])
}
