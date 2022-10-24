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

	arr = [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	fmt.Println(canFinish(4, arr))
}

func TestFindOrder(t *testing.T) {
	arr := [][]int{}
	fmt.Println(findOrder(1, arr))

	arr = [][]int{
		{1, 0},
	}
	fmt.Println(findOrder(2, arr))

	arr = [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	fmt.Println(findOrder(4, arr))
}
