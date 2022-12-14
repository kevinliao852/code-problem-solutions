func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := nums[0]
	n := nums[1]

	if n < m {
		n = m
	}

	for i := 2; i < len(nums); i++ {
		cur := nums[i]
		if i%2 == 0 {
			m += cur
			if m < n {
				m = n
			}
		} else {
			n += cur
			if n < m {
				n = m
			}
		}
	}

	var res int

	if n < m {
		res = m
	} else {
		res = n
	}
	return res
}
