func search(nums []int, target int) int {
	// binary search
	if nums[0] == target {
		return 0
	}

	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	var m int

	for i, j := 0, len(nums)-1; i <= j; {

		m = (i + j) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}

	}

	return -1
}
