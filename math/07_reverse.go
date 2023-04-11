package math

import "math"

// 难点在于判断是否溢出
func reverse(x int) int {
	isNegate := false
	if x < 0 {
		isNegate = true
		x = -x
	}
	res := 0
	for x != 0 {
		tmp := x%10
		if !isNegate &&
			(res > math.MaxInt32/10 ||
				(res == math.MaxInt32/10 && tmp > 7)) {
			return 0
		}
		if isNegate &&
			(res > -math.MinInt32/10 ||
				(res == -math.MinInt32 && tmp > 8)) {
			return 0
		}
		res = 10 * res + tmp
		x /= 10
	}
	if isNegate {
		return -res
	}
	return res
}
