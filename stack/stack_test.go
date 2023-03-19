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
	var rpn []string
	rpn = []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(rpn))

	rpn = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(rpn))

	rpn = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(rpn))
}

func TestMyQueue(t *testing.T)  {
	s := MyQueueConstructor()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}

func TestMinStack(t *testing.T)  {
	stack := Constructor()
	stack.Push(0)
	stack.Push(1)
	stack.Push(0)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}