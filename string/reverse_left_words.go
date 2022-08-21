package string

// 不申请额外空间，左移前k个字符
func reverseLeftWords(s string, k int) string {
	ss := []byte(s)
	reverseLeftWordsReverse(ss, 0, len(ss) - 1)
	reverseLeftWordsReverse(ss, 0, len(ss) - 1 - k)
	reverseLeftWordsReverse(ss, len(ss) - k, len(ss) - 1)
	return string(ss)
}

func reverseLeftWordsReverse(s []byte, left int, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
