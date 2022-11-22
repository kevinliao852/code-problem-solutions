func decodeMessage(key string, message string) string {
	hashmap := make(map[string]string)
	hashmap[" "] = " "

	set := "abcdefghijklmnopqrstuvwxyz"
	count := 0

	for i := 0; i < len(key); i++ {
		key_str := string(key[i])
		is_not_space := key_str != " "
		if _, ok := hashmap[key_str]; ok == false && is_not_space {
			hashmap[key_str] = string(set[count])
			count++
		}

	}

	l := make([]string, 0, 0)

	for i := 0; i < len(message); i++ {
		ch := hashmap[string(message[i])]
		l = append(l, ch)
	}

	return strings.Join(l, "")
} 
