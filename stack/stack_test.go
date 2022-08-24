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