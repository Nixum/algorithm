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