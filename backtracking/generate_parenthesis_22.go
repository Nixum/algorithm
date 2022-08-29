package backtracking

func generateParenthesis(n int) []string {
	generateParenthesisRes = make([]string, 0)
	BTWhenGenerateParenthesis(n, n, []byte{})
	return generateParenthesisRes
}

var generateParenthesisRes []string

func BTWhenGenerateParenthesis(left int, right int, track []byte) {
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		generateParenthesisRes = append(generateParenthesisRes, string(track))
		return
	}
	track = append(track, '(')
	BTWhenGenerateParenthesis(left - 1, right, track)
	track = track[:len(track) - 1]

	track = append(track, ')')
	BTWhenGenerateParenthesis(left, right - 1, track)
	track = track[:len(track) - 1]
}
