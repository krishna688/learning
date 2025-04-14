package main

import "fmt"

func main() {
	fmt.Println("working")

	fmt.Println(anagrams("abcd"))
}

func anagrams(input string) []string {

	if len(input) == 1 {
		fmt.Println("base condition")
		return []string{string(input[0])}
	}
	collections := []string{}

	ans := anagrams(input[1:])

	for _, an := range ans {
		for i := 0; i <= len(an); i++ {
			a := an[:i] + string(input[0]) + an[i:]

			collections = append(collections, a)
		}
	}

	return collections
}
