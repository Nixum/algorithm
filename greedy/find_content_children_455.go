package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 || len(g) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	sIndex := 0
	res := 0
	for i := 0; i < len(g); {
		if s[sIndex] >= g[i] {
			res++
			i++
		}
		sIndex++
		if sIndex >= len(s) {
			break
		}
	}
	return res
}

func findContentChildren2(g []int, s []int) int {
	if len(s) == 0 || len(g) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	gIndex := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if gIndex < len(g) && g[gIndex] <= s[i] {
			res++
			gIndex++
		}
	}
	return res
}
