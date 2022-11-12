func minimumSum(num int) int {
	// 1000 <= num <= 9999

	// Split num into two new integers new1 and new2
	// by using the digits
	// Leading zeros are allowed
	// and all the digits found in num **must be used**
	var new1 int
	var new2 int
	var sum int

	// int to slice
	num_str := strconv.Itoa(num)

	var num_slice []int

	for i := 0; i < len(num_str); i++ {
		str := string(num_str[i])
		n64, _ := strconv.ParseInt(str, 10, 32)
		n32 := int(n64)
		num_slice = append(num_slice, n32)
	}

	// sort
	sort.Ints(num_slice)

	// divide by 2

	new1 = num_slice[0]*10 + num_slice[2]
	new2 = num_slice[1]*10 + num_slice[3]

	sum = new1 + new2

	// return the minimum possible sum of new1 and new2.
	return sum
}
