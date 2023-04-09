func threeSum(nums []int) [][]int {
	var triplets [][]int

	// sort the list first
	bubbleSort(&nums)

	// let find the answer
	// make sure i, j, k are all unique

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// found
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else {
				if sum > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}

	return triplets
}

func bubbleSort(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		isSorted := true
		for j := 0; j < len(*nums)-i-1; j++ {
			if (*nums)[j] > (*nums)[j+1] {
				(*nums)[j], (*nums)[j+1] = (*nums)[j+1], (*nums)[j]
				isSorted = false
			}
		}

		if isSorted == true {
			break
		}
	}
}
