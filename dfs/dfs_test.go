package dfs

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	var arr [][]byte
	//arr = [][]byte{
	//	{'O'},
	//}
	//solve(arr)
	//fmt.Println(arr)

	arr = [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	solve(arr)
	fmt.Println(arr)
}