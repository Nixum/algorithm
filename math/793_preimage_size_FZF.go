package math

import "math"

func preimageSizeFZF(k int) int {
	if k <= 1 {
		return 5
	}
	return searchInPreimageSizeFZF(k) - searchInPreimageSizeFZF(k - 1)
}

// 返回当前数字的阶乘尾部有多少个零
func trailingZeroInPreimageSizeFZF(n int) int {
	res := 0
	for n != 0 {
		res += n/5
		n /= 5
	}
	return res
}

func searchInPreimageSizeFZF(k int) int {
	left := 0
	right := math.MaxInt64
	for left < right {
		mid := left + (right - left) / 2
		res := trailingZeroInPreimageSizeFZF(mid)
		if res > k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
