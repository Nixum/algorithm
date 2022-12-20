package slide_window

import "algorithm/common"

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
		// 保证窗口内无重复字符
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		max = maxN(max, right - left)
	}
	return max
}

// 错误的思路，这种方法无法解决从头到尾无重复的case，因为里面必须得出现重复才能得到答案
func lengthOfLongestSubstring2(s string) int {
	set := make(map[byte]int)
	left := 0
	right := 0
	res := 0
	for right < len(s) {
		c := s[right]
		if preIndex, exist := set[c]; exist {
			res = common.Max(res, right - left)
			left = common.Max(left, preIndex + 1)
		}
		set[c] = right
		right++
	}
	return res
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