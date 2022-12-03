func halvesAreAlike(s string) bool {
	var count int
	hashmap := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}

	for i, j := 0, len(s)/2; i < len(s)/2; i++ {
		if hashmap[string(s[i])] == true {
			count++
		}

		if hashmap[string(s[j])] == true {
			count--
		}
		j++
	}

	if len(s)%2 == 1 {
		if hashmap[string(s[len(s)-1])] == true {
			count--
		}
	}

	return count == 0
}
