package slide_window

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, len(s1))
	win := make(map[byte]int, len(s1))
	for i := range s1 {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if cn, exist := need[c]; exist {
			win[c]++
			if win[c] == cn {
				valid++
			}
		}
		for right - left >= len(s1) {
			fmt.Println(right, left, valid)
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if dn, exist := need[d]; exist {
				if dn == win[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return false
}
