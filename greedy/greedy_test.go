package greedy

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	var arr1 []int
	var arr2 []int
	arr1 = []int{1,2,3}
	arr2 = []int{1,1}
	fmt.Println(findContentChildren(arr1, arr2))

	arr1 = []int{1,2}
	arr2 = []int{1,2,3}
	fmt.Println(findContentChildren(arr1, arr2))
}
