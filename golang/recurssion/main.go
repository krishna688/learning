package main

import "fmt"

func main() {

	// slice := make([]int, 3, 5)

	// log.Print(len(slice))
	// log.Println(cap(slice))
	// slice2 := make([]int, 12, 13)
	// slice2 = append(slice2, slice...)
	// log.Print(len(slice2))
	// log.Println(cap(slice2))
	input := "abc"

	output := getAnagram(input)

	fmt.Println(output)
}
