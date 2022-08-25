package stack

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		return convert2IntInEvalRPN(tokens[0])
	}
	stack := make([]int, 0)
	res := 0
	for i := 0; i < len(tokens); i++ {
		if isNumberInEvalRPN(tokens[i]) {
			stack = append(stack, convert2IntInEvalRPN(tokens[i]))
			continue
		}
		lastIndex := len(stack) - 1
		switch tokens[i] {
		case "+":
			res = stack[lastIndex] + stack[lastIndex - 1]
		case "-":
			res = stack[lastIndex - 1] - stack[lastIndex]
		case "*":
			res = stack[lastIndex] * stack[lastIndex - 1]
		case "/":
			res = stack[lastIndex - 1] / stack[lastIndex]
		}
		stack = stack[:lastIndex - 1]
		stack = append(stack, res)
	}
	return res
}

func isNumberInEvalRPN(str string) bool {
	return str != "+" && str != "-" && str != "*" && str != "/"
}

func convert2IntInEvalRPN(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
