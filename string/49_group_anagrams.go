package string

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	sort2Val := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		byteStr := []byte(strs[i])
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})
		sort2Val[string(byteStr)] = append(sort2Val[string(byteStr)], strs[i])
	}
	for _, val := range sort2Val {
		res = append(res, val)
	}
	return res
}
