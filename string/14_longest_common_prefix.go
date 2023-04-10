package string

func longestCommonPrefix(strs []string) string {
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		tmp := strs[i]
		j := 0
		for ; j < len(tmp) && j < len(str); j++ {
			if str[j] != tmp[j] {
				break
			}
		}
		str = str[:j]
		if str == "" {
			return ""
		}
	}
	return str
}
