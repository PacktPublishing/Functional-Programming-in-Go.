package pkg

func IterativeFact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func RecursiveFact(n int) int {
	if n == 0 {
		return 1
	}
	return n * RecursiveFact(n-1)
}
