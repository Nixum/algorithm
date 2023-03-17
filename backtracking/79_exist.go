package backtracking

func exist(board [][]byte, word string) bool {
	strInExist = ""
	isVisit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		isVisit[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtrackingInExist(board, isVisit, i, j, word) {
				return true
			}
		}
	}
	return false
}

func backtrackingInExist(board [][]byte, isVisit [][]bool, r, c int, word string) bool {
	if !isInBoard(len(board), len(board[0]), r, c) {
		return false
	}
	if isVisit[r][c] {
		return false
	}
	str := strInExist + string(board[r][c])
	if str != word[:len(str)] {
		return false
	}
	if str == word {
		return true
	}
	strInExist += string(board[r][c])
	isVisit[r][c] = true
	for k := 0; k < 4; k++ {
		rr := r
		cc := c
		switch k {
		case 0:
			cc++
		case 1:
			rr++
		case 2:
			cc--
		case 3:
			rr--
		}
		if backtrackingInExist(board, isVisit, rr, cc, word) {
			return true
		}
	}
	strInExist = strInExist[:len(strInExist)-1]
	isVisit[r][c] = false
	return false
}

func isInBoard(row, col int, i, j int) bool {
	if i >= row || i < 0 ||
		j >= col || j < 0 {
		return false
	}
	return true
}

// 总体思路是对的，但是回溯的逻辑有问题，在字符串追加和方向回溯上有冲突
func exist2(board [][]byte, word string) bool {
	strInExist = ""
	resInExist = false
	isView := make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		isView[i] = make([]byte, len(board[i]))
	}
	backtrackingInExist2(board, isView, 0, 0, word)
	return resInExist
}

var strInExist string
var resInExist bool
func backtrackingInExist2(board [][]byte, isView [][]byte, r, c int, word string) {
	if r >= len(board) || r < 0 || c >= len(board[0]) || c < 0 {
		return
	}
	if resInExist  || len(strInExist) > len(word) {
		return
	}
	if strInExist != word[:len(strInExist)] {
		return
	}
	if strInExist == word {
		resInExist = true
		return
	}
	// 这里还有个问题是，递归进来之后，又是两层循环
	for i := r; i < len(board); i++ {
		for j := c; j < len(board[i]); j++ {
			if isView[i][j] == 1 {
				return
			}
			strInExist += string(board[i][j])
			if strInExist != word[:len(strInExist)] {
				strInExist = strInExist[:len(strInExist)-1]
				isView[i][j] = 0
				return
			}
			if strInExist == word {
				resInExist = true
				return
			}
			isView[i][j] = 1
			for k := 0; k < 4; k++ {
				rr := i
				cc := j
				switch k {
				case 0:
					cc++
				case 1:
					rr++
				case 2:
					cc--
				case 3:
					rr--
				}
				// 这里没有返回值，导致如果找到了，也不会立即返回
				backtrackingInExist2(board, isView, rr, cc, word)
			}
			strInExist = strInExist[:len(strInExist)-1]
			isView[i][j] = 0
		}
	}
}