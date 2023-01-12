package math

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j <= i; j++ {
			s1 := 0
			s2 := 0
			if j - 1 >= 0 {
				s1 = res[i-1][j-1]
			}
			if j < len(res[i-1]) {
				s2 = res[i-1][j]
			}
			row = append(row, s1 + s2)
		}
		res = append(res, row)
	}
	return res
}
