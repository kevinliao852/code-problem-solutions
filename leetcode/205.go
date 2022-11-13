func isIsomorphic(s string, t string) bool {
	hashmap := make(map[string]string)
	hashmap2 := make(map[string]string)
	// t.length == s.length

	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if v, ok := hashmap[ch]; ok != false {
			if v != string(t[i]) {
				return false
			}
		}

		ch2 := string(t[i])
		if v, ok := hashmap2[ch2]; ok != false {
			if v != string(s[i]) {
				return false
			}
		}

		hashmap[ch] = string(t[i])
		hashmap2[ch2] = string(s[i])
	}

	for i := 0; i < len(t); i++ {

	}

	return true
}
