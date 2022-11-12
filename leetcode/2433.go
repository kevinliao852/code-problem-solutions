func findArray(pref []int) []int {
	// xor
	// 1 1 -> 0
	// 0 1 -> 1
	// 1 0 -> 1
	// 0 0 -> 0

	var res []int

	for i := 0; i < len(pref); i++ {
		if i == 0 {
			res = append(res, pref[0])
			continue
		}

		r := pref[i-1] ^ pref[i]
		res = append(res, r)

	}

	return res
}
