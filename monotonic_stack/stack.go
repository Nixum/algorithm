package monotonic_stack

func top(monotonicStack []int) int {
	if len(monotonicStack) - 1  < 0 {
		return -1
	}
	return monotonicStack[len(monotonicStack) - 1]
}

func pop(monotonicStack []int) ([]int, int) {
	t := monotonicStack[len(monotonicStack) - 1]
	monotonicStack = monotonicStack[:len(monotonicStack) - 1]
	return monotonicStack, t
}
