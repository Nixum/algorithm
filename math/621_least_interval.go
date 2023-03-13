package math

import "algorithm/common"

// 思路：
// 如果总任务数少于下面计算出来的数量
// 重点是统计出任务种类最多的个数，因为(其个数-1) 乘 (任务时间间隔n+1)就能基本确认整个排列，
// 其他少于这个数量的任务就能插入其中，最后再插入剩余的任务(即任务个数等于最大任务数的任务数量)
// 如果总任务数大于上面计算出来的数量
// 多出来的任务插入到每个(n+1)后面即可，此时最短时间等于最大任务数
func leastInterval(tasks []byte, n int) int {
	charset := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		charset[tasks[i]-'A']++
	}
	taskNumMax := 0
	commonMaxTaskNum := 0
	for i := 0; i < 26; i++ {
		taskNumMax = common.Max(taskNumMax, charset[i])
	}
	for i := 0; i < 26; i++ {
		if charset[i] == taskNumMax {
			commonMaxTaskNum++
		}
	}
	if len(tasks) > (taskNumMax-1) * (n+1) + commonMaxTaskNum {
		return len(tasks)
	}
	return (taskNumMax-1) * (n+1) + commonMaxTaskNum
}
