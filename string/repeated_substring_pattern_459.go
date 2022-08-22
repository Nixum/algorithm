package string

// KMP算法
func repeatedSubstringPattern(s string) bool {
	next := buildNextInRepeatedSubstringPattern(s)
	size := len(s)
	// 如果是由可重复子串构成，其next数组的值在开始构成部分必定是升序的
	// 此时next[size - 1]就是倒数第二个重复子串的最后一个字符的下标
	// size - (next[size - 1]) 就是重复子串的个数
	// 所以如果取模能除尽，说明可构成
	return next[size - 1] != 0 &&
		size % (size - next[size - 1]) == 0
}

func buildNextInRepeatedSubstringPattern(s string) []int {
	next := make([]int, len(s))
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j - 1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
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
			// 优化，不对子串进行拼接，直接在父串上分割比较即可
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
		// 如果没有这个判断就超时了, 主要是判断子串是否能构成它本身，如果有余数说明不能构成
		if len(s) % i != 0 {
			continue
		}
		// 分割子串
		subStr := s[:i]
		str := ""
		// 拼接子串，当长度与父串相等时，进行比较
		for j := 1; j < len(s) / i + 1; j++ {
			str += subStr
			if str == s {
				return true
			}
		}
	}
	return false
}
