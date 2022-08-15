package hash

func intersection(a1 []int, a2 []int) []int {
	all := make([]int, 1001)
	for _, i := range a1 {
		all[i] = 1
	}
	res := make([]int, 0)
	for _, i := range a2 {
		if all[i] == 1 {
			all[i]++
			res = append(res, i)
		}
	}
	return res
}

func intersection2(a1 []int, a2 []int) []int {
	res := make([]int, 0)
	tmp := make(map[int]int)
	for i := 0; i < len(a1); i++ {
		tmp[a1[i]] = 1
	}
	for i := 0; i < len(a2); i++ {
		n := tmp[a2[i]]
		if n - 1 >= 0 {
			tmp[a2[i]]--
			res = append(res, a2[i])
		}
	}
	return res
}
