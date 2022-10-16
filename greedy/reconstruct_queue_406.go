package greedy

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		// 再按人数排
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 优先按身高排
		return people[i][0] > people[j][0]
	})
	// 以人数多少为下标进行插入
	res := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		res = append(res, people[i])
		// 往后挪位置，腾出people[i][1]
		copy(res[people[i][1] + 1 : ], res[people[i][1] : len(res) - 1])
		res[people[i][1]] = people[i]
	}
	// 简化写法
	//for _, p := range people {
	//	res = append(res, p)
	//	copy(res[p[1] + 1 : ], res[p[1] : ])
	//	res[p[1]] = p
	//}
	return res
}