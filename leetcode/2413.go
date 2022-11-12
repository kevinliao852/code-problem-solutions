func smallestEvenMultiple(n int) int {
	if n == 1 {
		return 2
	}

	r := n % 2

	if r == 1 {
		return n * 2
	}

	return n
}
