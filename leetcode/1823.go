func findTheWinner(n int, k int) int {
	var q []int

	for i := 1; i <= n; i++ {
		q = append(q, i)
	}

	var i int = 1

	for len(q) > 1 {

		e := q[0]
		q = q[1:]
		if i%k != 0 {
			q = append(q, e)
		}
		i++
	}

	return q[0]
}
