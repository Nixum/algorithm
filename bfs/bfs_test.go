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