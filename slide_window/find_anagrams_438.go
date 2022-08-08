package slide_window

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, len(p))
	win := make(map[byte]int, len(p))
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	valid := 0
	left := 0
	right := 0
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, exist := need[c]; exist {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}
		for right - left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, exist := need[d]; exist {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return res
}
