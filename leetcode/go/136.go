func singleNumber(nums []int) int {
	var res int

	for n := range nums {
		res ^= nums[n]
	}

	return res
}
