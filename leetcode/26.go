func removeDuplicates(nums []int) int {
	// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
	var pos int = 0
	var temp int = -1010
	for i := 0; i < len(nums); i++ {
		if temp != nums[i] {
			temp = nums[i]
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos
}
