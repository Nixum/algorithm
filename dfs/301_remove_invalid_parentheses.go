package dfs

func removeInvalidParentheses(s string) []string {
	deleteL := 0
	deleteR := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			deleteL++
		} else if s[i] == ')' {
			if deleteL == 0 {
				deleteR++
			} else {
				deleteL--
			}
		}
	}
	if deleteL == 0 && deleteR == 0 {
		return []string{s}
	}
	set := make(map[string]int)
	dfsInRemoveInvalidParentheses(s, 0, "", set, deleteL, deleteR, 0)
	res := make([]string, 0, len(set))
	for key := range set {
		res = append(res, key)
	}
	return res
}

// 剪枝版, diff表示左括号比右括号的差值，解决右括号比左括号多的场景，不需要再往下递归
func dfsInRemoveInvalidParentheses(s string, i int, res string, set map[string]int, deleteL, deleteR int, diff int) {
	if i == len(s) {
		if diff == 0 && isValidParentheses(res) {
			set[res] = 1
		}
		return
	}
	switch s[i] {
	case '(':
		// 要（
		res += string(s[i])
		dfsInRemoveInvalidParentheses(s, i+1, res, set, deleteL, deleteR, diff+1)
		res = res[:len(res)-1]
		// 不要（
		if deleteL > 0 {
			dfsInRemoveInvalidParentheses(s, i+1, res, set, deleteL-1, deleteR, 0)
		}
	case ')':
		// diff大于0才需要）
		if diff > 0 {
			res += string(s[i])
			dfsInRemoveInvalidParentheses(s, i+1, res, set, deleteL, deleteR, diff-1)
			res = res[:len(res)-1]
		}
		// 不要）
		if deleteR > 0 {
			dfsInRemoveInvalidParentheses(s, i+1, res, set, deleteL, deleteR-1, 0)
		}
	default:
		res += string(s[i])
		dfsInRemoveInvalidParentheses(s, i+1, res, set, deleteL, deleteR, 0)
		res = res[:len(res)-1]
	}
}

// 普通版
func dfsInRemoveInvalidParentheses1(s string, i int, res string, set map[string]int, deleteL, deleteR int) {
	if i == len(s) {
		if deleteR == 0 && deleteL == 0 && isValidParentheses(res) {
			set[res] = 1
		}
		return
	}
	switch s[i] {
	case '(':
		// 要（
		res += string(s[i])
		dfsInRemoveInvalidParentheses1(s, i+1, res, set, deleteL, deleteR)
		res = res[:len(res)-1]
		// 不要（
		if deleteL > 0 {
			dfsInRemoveInvalidParentheses1(s, i+1, res, set, deleteL-1, deleteR)
		}
	case ')':
		// 要）
		res += string(s[i])
		dfsInRemoveInvalidParentheses1(s, i+1, res, set, deleteL, deleteR)
		res = res[:len(res)-1]
		// 不要）
		if deleteR > 0 {
			dfsInRemoveInvalidParentheses1(s, i+1, res, set, deleteL, deleteR-1)
		}
	default:
		res += string(s[i])
		dfsInRemoveInvalidParentheses1(s, i+1, res, set, deleteL, deleteR)
		// 字母也要回退，才能回溯字母前的括号
		res = res[:len(res)-1]
	}
}

func isValidParentheses(p string) bool {
	left := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '(' {
			left++
		} else if p[i] == ')' {
			if left == 0 {
				return false
			} else {
				left--
			}
		}
	}
	return left == 0
}