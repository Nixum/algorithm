package string

func firstUniqChar(s string) int {
	bs := make([]int, 26)
	for i := 0; i < len(s); i++ {
		bs[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if bs[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
