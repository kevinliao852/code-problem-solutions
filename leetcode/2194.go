func cellsInRange(s string) []string {
	l := strings.Split(s, ":")
	a := l[0]
	b := l[1]

	iLen := int(b[0] - a[0])
	jLen := int(b[1] - a[1])

	var res []string

	for i := 0; i <= iLen; i++ {
		for j := 0; j <= jLen; j++ {
			str := string(a[0]+byte(i)) + string(a[1]+byte(j))
			res = append(res, str)
		}
	}

	return res
}
