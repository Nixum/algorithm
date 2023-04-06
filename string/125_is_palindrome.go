package string

import "strings"

func isPalindrome(s string) bool {
	left := 0
	right := len(s)-1
	for left < right {
		if !isAlphabet(s[left]) {
			left++
			continue
		}
		if !isAlphabet(s[right]) {
			right--
			continue
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphabet(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}