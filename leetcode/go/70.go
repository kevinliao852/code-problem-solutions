func climbStairs(n int) int {
	s := make([]int, n+1)

	return calFib(n, &s)
}

func calFib(n int, s *[]int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if (*s)[n] != 0 {
		return (*s)[n]
	}

	(*s)[n] = calFib(n-1, s) + calFib(n-2, s)

	return (*s)[n]
}
