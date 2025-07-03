package main

func noofways(n int) int {

	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}

	return noofways(n-1) + noofways(n-2) + noofways(n-3)
}
