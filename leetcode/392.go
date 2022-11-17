func isSubsequence(s string, t string) bool {
	var cnt int
	var res bool

	if s == "" {
		return true
	}

	for i := 0; i < len(t); i++ {
		if s[cnt] == t[i] {
			cnt++
			if cnt >= len(s) {
				break
			}
		}
	}

	res = len(s) == cnt
	return res
}
