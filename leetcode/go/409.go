func longestPalindrome(s string) int {
	var f bool
	var res int
	hashmap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		hashmap[string(s[i])] += 1
	}

	for _, v := range hashmap {
		if v == 1 {
			f = true
			continue
		}

		if v%2 == 1 {
			f = true
			res -= 1
		}

		res += v
	}

	if f == true {
		res += 1
	}

	return res
}
