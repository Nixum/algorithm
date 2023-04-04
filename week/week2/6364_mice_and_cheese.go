package week2

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	res := 0
	// 先假定所有奶酪都由2号鼠吃掉
	for i := 0; i < len(reward2); i++ {
		res += reward2[i]
	}
	// 计算reward1和reward2的比分差值，并进行排序
	dif := make([]int, len(reward2))
	for i := 0; i < len(reward2); i++ {
		dif[i] = reward1[i] - reward2[i]
	}
	sort.Ints(dif)
	// 取差值的前k个进行计算
	for i := 0; i < k; i++ {
		res += dif[len(dif)-i-1]
	}
	return res
}

// 错误的思路，无法解决 reward1和reward2 元素最优选择问题
func miceAndCheese2(reward1 []int, reward2 []int, k int) int {
	res := 0
	reward1Sort := make([][]int, 0, len(reward1))
	for i := 0; i < len(reward1); i++ {
		reward1Sort = append(reward1Sort, []int{i, reward1[i]})
	}
	sort.Slice(reward1Sort, func(i, j int) bool {
		return reward1Sort[i][1] > reward1Sort[j][1]
	})
	kNums := make(map[int]bool)
	for i := 0; i < k; i++ {
		res += reward1Sort[i][1]
		kNums[reward1Sort[i][0]] = true
	}
	for i := 0; i < len(reward2); i++ {
		if !kNums[i] {
			res += reward2[i]
		}
	}
	return res
}
