func isStrictlyPalindromic(n int) bool {

	for i := 2; i <= n-2; i++ {
		// check if each number is palindromic or not
		b := fmt.Sprintf("%b", i)
		for j, k := 0, len(b)-1; j < k; j++ {
			ch_start := string(b[j])
			ch_end := string(b[k])

			if ch_start != ch_end {
				return false
			}

			k--
		}
	}

	return true
}
