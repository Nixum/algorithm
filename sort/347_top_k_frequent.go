package sort

import sort2 "sort"

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}
	list := make([][]int, 0)
	for n, freq := range hash {
		list = append(list, []int{n, freq})
	}
	sort2.Slice(list, func(i, j int) bool {
		return list[i][1] > list[j][1]
	})
	for _, r := range list {
		res = append(res, r[0])
		k--
		if k == 0 {
			break
		}
	}
	return res
}
