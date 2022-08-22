package string

// 判断needle是否是haystack的子串，是的话就返回第一个索引下标
// KMP
// 理解什么是最长相同前后缀
// 比如 a a b a a f
//      0 1 2 3 4 5
// 每个下标对应的最长相同前后缀是
//      0 1 0 1 2 0
// 以 a a b a a b a a f a 文本串
//    0 1 2 3 4 5
//    a a b a a f         模式串
//    0 1 0 1 2 0
// 当找到下标5时，字符串和模式串不匹配，此时其前一个下标4的最长前后缀的值为2，
// 此时从模式串下标为2的地方开始继续匹配
// 最长相同前后缀统一减一即可得到next数组，比如值为
// -1 0 -1 0 1 -1，使用next数组和最长相同前后缀数组都可以实现KMP算法，只是有一点不同而已
func strStr(haystack string, needle string) int {
	next := buildNextInStrStr(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j - 1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

// 前缀：不包含最后一个字符的所有以第一个字符开头的连续子串
// 后缀：不包含第一个字符的所有以最后一个字符结尾的连续子串
// 前后缀的值的含义：下标i之前，包括i的字符串中，有多少长度相同的前后缀
func buildNextInStrStr(s string) []int {
	next := make([]int, len(s))
	// j 代表 前缀末尾的位置，也代表i以及i之前这个子串的最长相等前后缀的长度
	j := 0
	next[0] = j
	// i 代表后缀末尾的位置
	// 这里的目的是为了计算 前i个字符相同的前后缀的下标，然后存入next数组
	// i 总是大于 j
	for i := 1; i < len(s); i++ {
		// 前后缀不同，j向前回退
		for j > 0 && s[i] != s[j] {
			j = next[j - 1]
		}
		// 有相同的前后缀
		if s[i] == s[j] {
			j++
		}
		// 前缀长度赋给next[i]
		next[i] = j
	}
	return next
}

// 暴力解法
func strStr2(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			end := i + len(needle)
			if end > len(haystack) {
				end = len(haystack)
			}
			for haystack[i:end] == needle {
				return i
			}
		}
	}
	return -1
}