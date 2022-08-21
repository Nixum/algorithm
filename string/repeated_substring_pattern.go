package string

// KMP算法
func repeatedSubstringPattern(s string) bool {

	return false
}

// 对下面的优化, 本质上是减少字符串的拼接
func repeatedSubstringPattern2(s string) bool {
	max := len(s) / 2
	for i := 1; i <= max; i++ {
		// 如果没有这个判断就超时了
		if len(s) % i != 0 {
			continue
		}
		m := len(s) / i
		subStr := s[:i]
		j := 1
		for ; j < m; j++ {
			if subStr != s[j*i : i+j*i] {
				break
			}
		}
		if j == m {
			return true
		}
	}
	return false
}

// 暴力解法
// 由于子串不可能是由它本身构成，所以子串的长度范围为1~len/2
func repeatedSubstringPattern3(s string) bool {
	max := len(s) / 2
	for i := 1; i <= max; i++ {
		// 如果没有这个判断就超时了
		if len(s) % i != 0 {
			continue
		}
		subStr := s[:i]
		str := ""
		for j := 1; j < len(s) / i + 1; j++ {
			str += subStr
			if str == s {
				return true
			}
		}
	}
	return false
}
