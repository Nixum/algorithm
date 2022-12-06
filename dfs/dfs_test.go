package dfs

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	var arr [][]byte
	arr = [][]byte{
		{'O'},
	}
	solve(arr)
	fmt.Println(arr)

	arr = [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	solve(arr)
	fmt.Println(arr)
}

func TestMaxAreaOfIsland(t *testing.T) {
	var arr [][]int
	arr = [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	fmt.Println(maxAreaOfIsland(arr))

	arr = [][]int{
		{0,0,0,0,0,0,0,0},
	}
	fmt.Println(maxAreaOfIsland(arr))
}

func TestNumEnclaves(t *testing.T) {
	var arr [][]int
	arr = [][]int{
		{0,0,0,0},{1,0,1,0},{0,1,1,0},{0,0,0,0},
	}
	fmt.Println(numEnclaves(arr))

	arr = [][]int{
		{0,0,0,0,0,0,0,0},
	}
	fmt.Println(numEnclaves(arr))
}

func TestCountSubIslands(t *testing.T) {
	var arr1 [][]int
	var arr2 [][]int
	arr1 = [][]int{
		{1,1,1,0,0},
		{0,1,1,1,1},
		{0,0,0,0,0},
		{1,0,0,0,0},
		{1,1,0,1,1},
	}
	arr2 = [][]int{
		{1,1,1,0,0},
		{0,0,1,1,1},
		{0,1,0,0,0},
		{1,0,1,1,0},
		{0,1,0,1,0},
	}
	fmt.Println(countSubIslands(arr1, arr2))

	arr1 = [][]int{
		{1,1,1,1,0,0},
		{1,1,0,1,0,0},
		{1,0,0,1,1,1},
		{1,1,1,0,0,1},
		{1,1,1,1,1,0},
		{1,0,1,0,1,0},
		{0,1,1,1,0,1},
		{1,0,0,0,1,1},
		{1,0,0,0,1,0},
		{1,1,1,1,1,0},
	}
	arr2 = [][]int{
		{1,1,1,1,0,1},
		{0,0,1,0,1,0},
		{1,1,1,1,1,1},
		{0,1,1,1,1,1},
		{1,1,1,0,1,0},
		{0,1,1,1,1,1},
		{1,1,0,1,1,1},
		{1,0,0,1,0,1},
		{1,1,1,1,1,1},
		{1,0,0,1,0,0},
	}
	fmt.Println(countSubIslands(arr1, arr2))
}

func TestFindCheapestPrice(t *testing.T) {
	var edges = [][]int{}
	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice2(len(edges), edges, 0, 2, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	fmt.Println(findCheapestPrice2(len(edges), edges, 0, 3, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice2(len(edges), edges, 0, 2, 0))
}