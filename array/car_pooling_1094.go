package array

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		// 在trip[2]已经下车
		diff[trip[2]] -= trip[0]

	}
	res := make([]int, 1001)
	res[0] = diff[0]
	if res[0] > capacity {
		return false
	}
	for i := 1; i < len(diff); i++ {
		res[i] = res[i - 1] + diff[i]
		if res[i] > capacity {
			return false
		}
	}
	return true
}
