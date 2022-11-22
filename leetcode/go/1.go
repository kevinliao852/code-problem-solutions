func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var x int
	var y int

	for i := 0; i < len(nums); i++ {

		cv := nums[i] // current value

		if v, ok := m[target-cv]; ok {
			x = v
			y = i

			// only one solution
			break
		}
		m[cv] = i // store index
	}

	return []int{x, y}
}
