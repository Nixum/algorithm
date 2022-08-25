package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}

func TestEvalRPN(t *testing.T) {
	rpn := []string{}
	rpn = []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(rpn))

	rpn = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(rpn))

	rpn = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(rpn))
}