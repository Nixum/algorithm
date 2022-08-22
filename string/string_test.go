package string

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte("abcdefg")
	reverseString(s[0:2])
	fmt.Println(s, string(s))
	s = []byte("abcd")
	reverseString(s)
	fmt.Println(s, string(s))
}

func TestReverseStr(t *testing.T) {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("We are happy."))
	fmt.Println(reverseWords("  hello world!  "))
}

func TestReverseLeftWords(t *testing.T) {
	fmt.Println(reverseLeftWords("Wearehappy", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

func TestRepeatedSubstringPattern(t *testing.T) {
	fmt.Println(repeatedSubstringPattern2("abab"))
	fmt.Println(repeatedSubstringPattern2("aba"))
	fmt.Println(repeatedSubstringPattern2("abcabcabcabc"))
	fmt.Println()
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func TestStrStr(t *testing.T) {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}