package main

func getAnagram(input string) []string {

	if len(input) == 1 {
		return []string{string(input[0])}
	}

	collections := make([]string, 0)

	anagrams := getAnagram(input[1:])

	for _, anagram := range anagrams {
		for i := 0; i <= len(anagram); i++ {
			an := anagram[:i] + string(input[0]) + anagram[i:]
			collections = append(collections, an)
		}
	}

	return collections
}
