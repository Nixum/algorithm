package interval_problem

import "algorithm/common"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	fLen := len(firstList)
	sLen := len(secondList)
	fIndex := 0
	sIndex := 0
	for sIndex < sLen && fIndex < fLen {
		a1 := firstList[fIndex][0]
		a2 := firstList[fIndex][1]
		b1 := secondList[sIndex][0]
		b2 := secondList[sIndex][1]
		if b2 >= a1 && b1 <= a2 {
			res = append(res, []int{common.Max(a1, b1), common.Min(a2, b2)})
		}
		if b2 > a2 {
			fIndex++
		} else {
			sIndex++
		}
	}
	return res
}

// 解决下面的问题，但实际上只是方法1的繁琐写法罢了
func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	fLen := len(firstList)
	sLen := len(secondList)
	fIndex := 0
	sIndex := 0
	for sIndex < sLen && fIndex < fLen {
		a1 := firstList[fIndex][0]
		a2 := firstList[fIndex][1]
		b1 := secondList[sIndex][0]
		b2 := secondList[sIndex][1]
		if b2 >= a1 && b1 <= a2 {
			if a1 <= b1 && a2 < b2 {
				res = append(res, []int{b1, a2})
				fIndex++
			} else if a1 > b1 && a2 >= b2 {
				res = append(res, []int{a1, b2})
				sIndex++
			} else if a1 >= b1 && a2 > b2 {
				res = append(res, []int{a1, b2})
				sIndex++
			} else if a1 < b1 && a2 <= b2 {
				res = append(res, []int{b1, a2})
				fIndex++
			} else if a1 > b1 && a2 < b2 {
				res = append(res, []int{a1, a2})
				fIndex++
			} else if a1 < b1 && a2 > b2 {
				res = append(res, []int{b1, b2})
				sIndex++
			} else {
				res = append(res, []int{b1, b2})
				fIndex++
				sIndex++
			}
		} else {
			if b2 > a2 {
				fIndex++
			} else {
				sIndex++
			}
		}
	}
	return res
}

// 错误的思路，虽然是分情况了，但是没有分出区别，所以没取到交集
// 这种思路，在两个区间没有交集的时候就会出错
// 比如 {0, 2}, {5, 10}, {13, 23}, {24, 25},
//      {1, 5}, {8, 12}, {15, 24}, {25, 26},
// 会出现{13, 12}这种结果
// 还有就是会处理不了两个集合相等的问题
func intervalIntersection3(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	fLen := len(firstList)
	sLen := len(secondList)
	fIndex := 0
	sIndex := 0
	for sIndex < sLen && fIndex < fLen {
		a1 := firstList[fIndex][0]
		a2 := firstList[fIndex][1]
		b1 := secondList[sIndex][0]
		b2 := secondList[sIndex][1]
		if a1 <= b1 && a2 < b2 {
			res = append(res, []int{b1, a2})
			fIndex++
		} else if a1 > b1 && a2 >= b2 {
			res = append(res, []int{a1, b2})
			sIndex++
		} else if a1 >= b1 && a2 > b2 {
			res = append(res, []int{a1, b2})
			sIndex++
		} else if a1 < b1 && a2 <= b2 {
			res = append(res, []int{b1, a2})
			fIndex++
		} else if a1 > b1 && a2 < b2 {
			res = append(res, []int{a1, a2})
			fIndex++
		} else if a1 < b1 && a2 > b2 {
			res = append(res, []int{b1, b2})
			sIndex++
		}
	}
	return res
}