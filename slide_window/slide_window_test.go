package slide_window

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2,3,1,2,0,3}))

	fmt.Println(minSubArrayLen(4, []int{1,4,4}))

	fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1}))

	fmt.Println(minSubArrayLen(11, []int{1,2,3,4,5}))
}

func TestTotalFruit(t *testing.T) {

	fmt.Println(totalFruit([]int{1,2,1}))

	fmt.Println(totalFruit([]int{0,1,2,2}))

	fmt.Println(totalFruit([]int{1,2,3,2,2}))

	fmt.Println(totalFruit([]int{3,3,3,1,2,1,1,2,3,3,4}))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "b"))
	fmt.Println(minWindow("a", "aa"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring(""))
}

func TestFindAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func TestCheckInclusion(t *testing.T) {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("abcdxabcde", "abcdeabcdx"))
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{7,2,4}, 1))
	fmt.Println(maxSlidingWindow([]int{7,2,4}, 2))
}
