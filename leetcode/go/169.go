func majorityElement(nums []int) int {
	var candidate int
	var count int

	// Boyer-Moore Majority Vote Algorithm

	for i := range nums {
		n := nums[i]
		if count == 0 {
			candidate = n
		}

		if candidate == n {
			count++
		} else {
			count--
		}
	}

	return candidate
}
