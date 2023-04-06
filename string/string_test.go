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
	//fmt.Println(strStr("hello", "ll"))
	//fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("sadbutsad", "sad"))
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{""}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}

func TestLargestNumber(t *testing.T) {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))

	nums = []int{3,30,34,5,9}
	fmt.Println(largestNumber(nums))

	nums = []int{0, 0, 0}
	fmt.Println(largestNumber(nums))
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("-91283472332")) // -2147483648
	fmt.Println(myAtoi("9223372036854775808")) // 2147483647
	fmt.Println(myAtoi("+-12")) // 0
	fmt.Println(myAtoi("asdasd+1")) // 0
	fmt.Println(myAtoi("+1")) // 1
	fmt.Println(myAtoi("+42")) // 42
	fmt.Println(myAtoi("42")) // 42
	fmt.Println(myAtoi("   -42")) // -42
	fmt.Println(myAtoi("4193 with words")) // 4193
	fmt.Println(myAtoi("words and 987")) // 0
	fmt.Println(myAtoi("-91283472332")) // -2147483648
}

func TestDecodeString(t *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
	fmt.Println(decodeString("10[lc]"))
}

func TestIsPalindrome(t *testing.T) {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}