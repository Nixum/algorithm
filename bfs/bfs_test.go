package bfs

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	dead := []string{"0201","0101","0102","1212","2002"}
	fmt.Println(openLock(dead, "0202"))
}

func TestSlidingPuzzle(t *testing.T) {
	board := [][]int{
		{1,2,3},
		{4,0,5},
	}
	fmt.Println(slidingPuzzle(board))
}

func TestFindCheapestPrice(t *testing.T) {
	var edges = [][]int{}
	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 2, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 3, 1))

	edges = [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(len(edges), edges, 0, 2, 0))
}