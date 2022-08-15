package hash

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for i, v := range nums {
		remain := target - v
		if preIndex, exist := m[remain]; exist {
			return []int{preIndex, i}
		}
		m[v] = i
	}
	return res
}
