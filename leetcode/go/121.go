func maxProfit(prices []int) int {
	var max int
	var idx int

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			continue
		}

		if prices[idx] > prices[i] {
			idx = i
			continue
		}

		if max < prices[i]-prices[idx] {
			max = prices[i] - prices[idx]
		}
	}
	return max
}
