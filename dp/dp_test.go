package dp

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {
	var arr []int
	arr = []int{1,5,11,5}
	fmt.Println(canPartition(arr))

	arr = []int{1, 2,3,5}
	fmt.Println(canPartition(arr))

	arr = []int{14,9,8,4,3,2}
	fmt.Println(canPartition(arr))

	arr = []int{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97}
	fmt.Println(canPartition(arr))
}

func TestLastStoneWeightII(t *testing.T) {
	var arr []int
	arr = []int{2,7,4,1,8,1}
	fmt.Println(lastStoneWeightII(arr))

	arr = []int{31,26,33,21,40}
	fmt.Println(lastStoneWeightII(arr))
}

func TestFindMaxForm(t *testing.T) {
	var arr []string
	arr = []string{"10","0001","111001","1","0"}
	fmt.Println(findMaxForm(arr, 5, 3))
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func TestRobII(t *testing.T) {
	fmt.Println(robII([]int{2,3,2}))
	fmt.Println(robII([]int{1,2,3,1}))
	fmt.Println(robII([]int{0}))
}

func TestMaxProfitIII(t *testing.T) {
	fmt.Println(maxProfitIII([]int{3,3,5,0,0,3,1,4}))
	fmt.Println(maxProfitIII([]int{1,2,3,4,5}))
	fmt.Println(maxProfitIII([]int{7,6,4,3,1}))
}

func TestMaxProfitIV(t *testing.T) {
	fmt.Println(maxProfitIV(2, []int{2, 4, 1}))
	fmt.Println(maxProfitIV(2, []int{3, 2, 6, 5, 0, 3}))
}

func TestMaxProfitV(t *testing.T) {
	fmt.Println(maxProfitV([]int{1,2,3,0,2}))
	fmt.Println(maxProfitV([]int{1}))
	fmt.Println(maxProfitV([]int{1,2}))
}

func TestMaxProfitVI(t *testing.T) {
	fmt.Println(maxProfitVI2([]int{1,3,2,8,4,9}, 2))
	fmt.Println(maxProfitVI2([]int{1,3,7,5,10,3}, 3))
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7}))
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2,2}))
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray2([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("bbbab"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aaaaa"))
}

func TestMinimumDeleteSum(t *testing.T) {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func TestFindCheapestPrice(t *testing.T) {
	var edges = [][]int{}
	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 2, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 3, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 2, 0))
}

func TestSuperEggDrop(t *testing.T) {
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(3, 14))
}

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))

	nums = []int{-2,0,-1}
	fmt.Println(maxProduct(nums))
}
