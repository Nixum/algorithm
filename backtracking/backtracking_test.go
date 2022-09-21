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
