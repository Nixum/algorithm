package monotonic_stack

// 单调栈，思路：
// 栈存数组下标，方便存结果
// 判断栈顶是否小于当前温度，如果是，更新栈顶对应下标的结果，栈顶出栈，继续比较下一个栈顶
// 当前温度下标进栈
func dailyTemperatures(temperatures []int) []int {
	dQueue := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(dQueue) > 0 && temperatures[dQueue[len(dQueue)-1]] < temperatures[i] {
			preIndex := dQueue[len(dQueue)-1]
			dQueue = dQueue[:len(dQueue)-1]
			ans[preIndex] = i - preIndex
		}
		dQueue = append(dQueue, i)
	}
	return ans
}
