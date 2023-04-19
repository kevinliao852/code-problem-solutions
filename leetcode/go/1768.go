func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder

	i := 0

	for i < len(word1) && i < len(word2) {
		sb.WriteString(string(word1[i]))
		sb.WriteString(string(word2[i]))
		i++
	}

	for i < len(word1) {
		sb.WriteString(string(word1[i]))
		i++
	}

	for i < len(word2) {
		sb.WriteString(string(word2[i]))
		i++
	}

	return sb.String()
}
