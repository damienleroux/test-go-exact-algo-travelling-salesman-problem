package main

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
