func uniqueOccurrences(arr []int) bool {
	hashmap1 := make(map[int]int)
	hashmap2 := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		if _, ok := hashmap1[arr[i]]; ok != true {
			hashmap1[arr[i]] = 1
		} else {
			hashmap1[arr[i]]++
		}
	}

	for _, v := range hashmap1 {
		if _, ok := hashmap2[v]; ok != true {
			hashmap2[v] = true
		} else {
			return false
		}
	}

	return true
}
