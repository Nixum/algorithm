package graph

import (
	"fmt"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	arr := [][]int{
		{1,2},
		{3},
		{3},
		{},
	}
	fmt.Println(allPathsSourceTarget(arr))
}

func TestCanFinish(t *testing.T) {
	arr := [][]int{}
	fmt.Println(canFinish(2, arr))

	arr = [][]int{
		{1, 0},
	}
	fmt.Println(canFinish(2, arr))

	arr = [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(canFinish(2, arr))
}