package slide_window

import "math"

func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	if len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	valid := 0
	right := 0
	left := 0
	start := 0
	min := math.MaxInt64
	window := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		if _, exist := need[c]; exist {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if right - left < min {
				start = left
				min = right - left
			}
			d := s[left]
			left++
			if _, exist := need[d]; exist {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	str := ""
	for i := 0; i < min; i++ {
		str += string(s[start])
		start++
	}
	return str
}
