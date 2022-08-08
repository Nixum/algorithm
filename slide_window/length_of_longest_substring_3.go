package slide_window

// 最长不重复子串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	left := 0
	right := 0
	win := make(map[byte]int, 2)
	for right < len(s) {
		c := s[right]
		win[c]++
		right++
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		max = maxN(max, right - left)
	}
	return max
}


// 错误的思路
func wrongLengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	left := 0
	right := 0
	subStrMap := make(map[byte]int, 2)
	for right < len(s) {
		c := s[right]
		right++
		if _, exist := subStrMap[c]; !exist {
			subStrMap[c]++
		} else {
			targetLen := len(subStrMap) - 1
			for len(subStrMap) != targetLen {
				d := s[left]
				left++
				if _, exist := subStrMap[d]; exist {
					subStrMap[d]--
				}
				if subStrMap[d] == 0 {
					delete(subStrMap, d)
				}
			}
			subStrMap[c]++
		}
		max = maxN(max, len(subStrMap))
	}
	return max
}