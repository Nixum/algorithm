package stack

// 下面方法的简化版本
func isValid(s string) bool {
	m := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack) - 1] == m[s[i]] {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) - 1 < 0 {
			return false
		}
		lastChar := stack[len(stack) - 1]
		switch lastChar {
		case '{':
			if s[i] == '}' {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		case '[':
			if s[i] == ']' {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		case '(':
			if s[i] == ')' {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
