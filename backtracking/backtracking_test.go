package backtracking

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	//fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))
}

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func TestCombinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

func TestCombinationSum(t *testing.T) {
	var arr []int
	arr = []int{2,3,6,7}
	fmt.Println(combinationSum(arr, 7))
	arr = []int{2,3,5}
	fmt.Println(combinationSum(arr, 8))
	arr = []int{2}
	fmt.Println(combinationSum(arr, 1))
}

func TestCombinationSum2(t *testing.T) {
	var arr []int
	arr = []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum2(arr, 8))
	arr = []int{2,5,2,1,2}
	fmt.Println(combinationSum2(arr, 5))
	arr = []int{2}
	fmt.Println(combinationSum2(arr, 1))
	arr = []int{2, 2, 2}
	fmt.Println(combinationSum2(arr, 2))
}

func TestPartition(t *testing.T) {
	arr := "aab"
	fmt.Println(partition(arr))
	arr = "a"
	fmt.Println(partition(arr))
}

func TestRestoreIpAddresses(t *testing.T) {
	arr := "25525511135"
	fmt.Println(restoreIpAddresses(arr))
	arr = "0000"
	fmt.Println(restoreIpAddresses(arr))
	arr = "1111"
	fmt.Println(restoreIpAddresses(arr))
	arr = "010010"
	fmt.Println(restoreIpAddresses(arr))
	arr = "101023"
	fmt.Println(restoreIpAddresses(arr))
}
