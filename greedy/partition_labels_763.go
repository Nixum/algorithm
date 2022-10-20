package greedy

import "algorithm/common"

func partitionLabels(s string) []int {
	alphabet := make([]int, 27)
	for i := 0; i < len(s); i++ {
		alphabet[s[i] - 'a'] = i
	}
	res := make([]int, 0)
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		right = common.Max(right, alphabet[s[i] - 'a'])
		if i == right {
			res = append(res, right - left + 1)
			left = i + 1
		}
	}
	return res
}