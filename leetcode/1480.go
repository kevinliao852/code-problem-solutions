func runningSum(nums []int) []int {
	var res []int
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res = append(res, sum)
	}

	return res
}
