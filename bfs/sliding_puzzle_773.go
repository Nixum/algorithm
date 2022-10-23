package bfs

import "fmt"

func slidingPuzzle(board [][]int) int {
	target := "123450"
	neighbor := [][]int {
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	visited := make(map[string]int)
	str := convertInSlidingPuzzle(board)
	visited[str] = 1
	queue := make([]string, 0)
	queue = append(queue, str)
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			str, queue = boardStrQueuePop(queue)
			if str == "" {
				continue
			}
			if str == target {
				return res
			}
			zeroIndex := 0
			for index, c := range str {
				if c - '0' == 0 {
					zeroIndex = index
					break
				}
			}
			for _, nextIndex := range neighbor[zeroIndex] {
				newStr := swapInSlidingPuzzle(str, zeroIndex, nextIndex)
				if _, exist := visited[newStr]; !exist {
					queue = append(queue, newStr)
					visited[newStr] = 1
				}
			}
		}
		res++
	}
	return -1
}

func boardStrQueuePop(queue []string) (string, []string) {
	if len(queue) == 0 {
		return "", queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}

func swapInSlidingPuzzle(str string, zeroIndex, nextIndex int) string {
	b := []byte(str)
	b[zeroIndex], b[nextIndex] = b[nextIndex], b[zeroIndex]
	return string(b)
}

// 思路是对的，但是超时了
func slidingPuzzle2(board [][]int) int {
	visited := make(map[string]int)
	str := convertInSlidingPuzzle(board)
	visited[str] = 1
	m := len(board) - 1
	n := len(board[0]) - 1
	queue := make([][][]int, 0)
	queue = append(queue, board)
	res := 0
	var b [][]int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			b, queue = boardQueuePop(queue)
			if len(b) == 0 {
				continue
			}
			str := convertInSlidingPuzzle(b)
			if str == "123450" {
				return res
			}
			zeroIIndex := 0
			zeroJIndex := 0
			for iIndex, row := range b {
				for jIndex, col := range row {
					if col == 0 {
						zeroIIndex = iIndex
						zeroJIndex = jIndex
					}
				}
			}
			if zeroIIndex - 1 >= 0 && zeroJIndex - 1 >= 0 {
				nB := slideInSlidingPuzzle(board, zeroIIndex, zeroJIndex, zeroIIndex - 1, zeroJIndex - 1)
				if _, exist := visited[convertInSlidingPuzzle(nB)]; !exist {
					queue = append(queue, nB)
				}
			}
			if zeroIIndex + 1 <= m && zeroJIndex - 1 >= 0 {
				nB := slideInSlidingPuzzle(board, zeroIIndex, zeroJIndex, zeroIIndex + 1, zeroJIndex - 1)
				if _, exist := visited[convertInSlidingPuzzle(nB)]; !exist {
					queue = append(queue, nB)
				}
			}
			if zeroIIndex - 1 >= 0 && zeroJIndex + 1 <= n {
				nB := slideInSlidingPuzzle(board, zeroIIndex, zeroJIndex, zeroIIndex - 1, zeroJIndex + 1)
				if _, exist := visited[convertInSlidingPuzzle(nB)]; !exist {
					queue = append(queue, nB)
				}
			}
			if zeroIIndex + 1 <= m && zeroJIndex + 1 <= n {
				nB := slideInSlidingPuzzle(board, zeroIIndex, zeroJIndex, zeroIIndex + 1, zeroJIndex + 1)
				if _, exist := visited[convertInSlidingPuzzle(nB)]; !exist {
					queue = append(queue, nB)
				}
			}
		}
		res++
	}
	return -1
}

func convertInSlidingPuzzle(board [][]int) string {
	str := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			str += fmt.Sprintf("%d", board[i][j])
		}
	}
	return str
}

func slideInSlidingPuzzle(board [][]int, iIndex, jIndex int, nextIIndex, nextJIndex int) [][]int {
	newBoard := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		newBoard[i] = make([]int, len(board[i]))
	}
	copy(newBoard, board)
	newBoard[nextIIndex][nextJIndex], newBoard[iIndex][jIndex] = newBoard[iIndex][jIndex], newBoard[nextIIndex][nextJIndex]
	return newBoard
}

func boardQueuePop(queue [][][]int) ([][]int, [][][]int) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}
