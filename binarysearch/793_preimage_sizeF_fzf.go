package binarysearch

import (
	"math"
)

// 阶乘为数小于等于k个0的个数 - 阶乘为数小于等于(k-1)个0的个数
func preimageSizeFZF(k int) int {
	return preimageSizeFZFLTENum(k) - preimageSizeFZFLTENum(k -1)
}

func preimageSizeFZFTrailingZeros(n int) int {
	res := 0
	for i := 5; i <= n; i *= 5 {
		res += n / i
	}
	return res
}

// 该方法最终只能知道阶乘尾数小于或等于k个0的当前数字是多少
func preimageSizeFZFLTENum(k int) int {
	left := 0
	right := math.MaxInt64
	for left < right {
		mid := left + (right - left) / 2
		rk := preimageSizeFZFTrailingZeros(mid)
		if rk > k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
