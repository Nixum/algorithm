package array

// 差分数组解法
func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for _, u := range updates {
		// 因为下标是从0开始的，所以不用减
		diff[u[0]] += u[2]
		// 这一步比较关键
		if u[1] + 1 < length {
			diff[u[1] + 1] -= u[2]
		}
	}
	res := make([]int, length)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i - 1] + diff[i]
	}
	return res
}
