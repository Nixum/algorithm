package week2

import "sort"

func findMatrix(nums []int) [][]int {
	// 根据总类和次数进行排序
	// 按总类和次数填充二维数组
	hash := make(map[int]int, 0)
	for _, n := range nums {
		hash[n]++
	}
	array := make([][]int, 0, len(hash))
	for key, value := range hash {
		array = append(array, []int{key, value})
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i][1] > array[j][1]
	})
	matrix := make([][]int, array[0][1])
	row := 0
	for {
		for i := 0; i < len(array); i++ {
			if array[i][1] > 0 {
				matrix[row] = append(matrix[row], array[i][0])
				array[i][1]--
			}
		}
		row++
		if row >= len(matrix) {
			break
		}
	}
	return matrix
}
