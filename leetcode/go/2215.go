func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	var answer0 []int
	var answer1 []int

	for _, v := range nums1 {
		if _, ok := m1[v]; ok {
			m1[v] += 1
		} else {
			m1[v] = 1
		}
	}

	for _, v := range nums2 {
		if _, ok := m2[v]; ok {
			m2[v] += 1
		} else {
			m2[v] = 1
		}
	}

	for v, _ := range m1 {
		if _, ok := m2[v]; !ok {
			answer0 = append(answer0, v)
		}
	}

	for v, _ := range m2 {
		if _, ok := m1[v]; !ok {
			answer1 = append(answer1, v)
		}
	}

	return [][]int{answer0, answer1}
}
