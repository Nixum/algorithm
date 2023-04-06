package slide_window

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
	window := make(map[byte]int)
	valid := 0
	right := 0
	left := 0

	str := ""
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
			if len(str) == 0 || len(str) > (right-left) {
				str = s[left:right]
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
	return str
}
