package math

// 输入一个n，从0开始，把n以内的带7或者7的倍数打印出来
func find7Num(n int) []int {
	res := make([]int, 0)
	i := 0
	for i < n {
		if i % 7 == 0 {
			res = append(res, i)
		} else {
			tmp := i
			for tmp != 0 {
				if tmp % 10 == 7 {
					res = append(res, i)
					break
				}
				tmp /= 10
			}
		}
		i++
	}
	return res
}
