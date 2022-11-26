func minCostClimbingStairs(cost []int) int {

	temp := make([]int, len(cost))
	temp[0] = cost[0]
	temp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		if temp[i-1] > temp[i-2] {
			temp[i] = temp[i-2] + cost[i]
		} else {
			temp[i] = temp[i-1] + cost[i]
		}
	}

	if temp[len(cost)-1] > temp[len(cost)-2] {
		return temp[len(cost)-2]
	} else {
		return temp[len(cost)-1]
	}
}
