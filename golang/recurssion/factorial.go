package main

func Fact01(num int) int {
	product := 1
	for i := 2; i <= num; i++ {
		product *= i

	}
	return product
}

func Fact02(num int) int {
	if num == 1 {
		return 1
	}

	return num * Fact02(num-1)
}
