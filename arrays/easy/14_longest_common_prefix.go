package easy

func longestCommonPrefixHorizontalScanning(strs []string) string {

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0

		for j < min(len(prefix), len(strs[i])) {
			if prefix[j] != strs[i][j] {
				break
			}
			j++
		}
		prefix = prefix[:j]
	}
	return prefix
}

func longestCommonPrefixVerticalScanning(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
