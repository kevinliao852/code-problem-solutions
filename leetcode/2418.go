func sortPeople(names []string, heights []int) []string {
	hashmap := make(map[int]string)

	for i := 0; i < len(heights); i++ {
		hashmap[heights[i]] = names[i]
	}

	h := heights[:]

	sort.Sort(sort.Reverse(sort.IntSlice(h)))

	res := []string{}

	for i := 0; i < len(h); i++ {
		res = append(res, hashmap[h[i]])
	}

	return res
} 
