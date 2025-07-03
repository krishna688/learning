package main

func Fib01(input int) []int {

	if input == 1 {
		return []int{0}
	}
	out := []int{0, 1}
	for i := 2; i < input; i++ {
		out = append(out, out[i-1]+out[i-2])
	}

	return out
}
