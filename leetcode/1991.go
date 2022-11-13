func findMiddleIndex(nums []int) int {
	// 1 <= nums.length <= 104
	// -1000 <= nums[i] <= 1000
	var sum int
	var prefixSum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if prefixSum*2+nums[i] == sum {
			return i
		}
		prefixSum += nums[i]
	}

	return -1
}

// the same as leetcode 724
