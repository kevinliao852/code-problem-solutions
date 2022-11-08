func makeGood(s string) string {
	slen := len(s)
	stack := make([]string, 0, slen)

	for i := 0; i < slen; i++ {

		ch := string(s[i])

		if len(stack) == 0 {
			stack = append(stack, string(ch))
			continue
		}

		top := stack[len(stack)-1]
		isSameInUpper := strings.ToUpper(top) == strings.ToUpper(ch)
		isNotSame := top != ch

		if isSameInUpper && isNotSame {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return strings.Join(stack, "")
}
