package stack

func removeDuplicates(s string) string {
	if s == "" {
		return ""
	}
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		lastIndex := len(stack) - 1
		if len(stack) > 0 && s[i] == stack[lastIndex] {
			stack = stack[:lastIndex]
		} else {
			stack = append(stack, s[i])
		}

		// 不能这么写，因为如果len(stack) == 0, 会导致lastIndex越界
		//if len(stack) > 0 && s[i] != stack[lastIndex] {
		//	stack = append(stack, s[i])
		//} else {
		//	stack = stack[:lastIndex]
		//}
	}
	return string(stack)
}
