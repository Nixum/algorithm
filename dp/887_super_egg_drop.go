package dp

import (
	"algorithm/common"
	"math"
)

func superEggDrop(k int, n int) int {
	memo := make([][]int, k + 1)
	for i := 0; i <= k; i++ {
		memo[i] = make([]int, n + 1)
		for j := 0; j <= n; j++ {
			memo[i][j] = math.MaxInt64
		}
	}
	return dpInSuperEggDrop(k, n, memo)
}

func dpInSuperEggDrop(k int, n int, memo [][]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if memo[k][n] != math.MaxInt64 {
		return memo[k][n]
	}
	res := math.MaxInt64
	left, right := 1, n
	for left <= right {
		mid := left + (right - left) / 2
		broken := dpInSuperEggDrop(k - 1, mid - 1, memo)
		notBroken := dpInSuperEggDrop(k, n - mid, memo)
		if broken > notBroken {
			right = mid - 1
			res = common.Min(res, broken + 1)
		} else {
			left = mid + 1
			res = common.Min(res, notBroken + 1)
		}
	}
	memo[k][n] = res
	return res
}

// 二分搜索，思路是对的，但是会超时，需要加入备忘录
func dpInSuperEggDrop2(k int, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	res := math.MaxInt64
	left, right := 1, n
	for left <= right {
		mid := left + (right - left) / 2
		broken := dpInSuperEggDrop2(k - 1, mid - 1)
		notBroken := dpInSuperEggDrop2(k, n - mid)
		if broken > notBroken {
			right = mid - 1
			res = common.Min(res, broken + 1)
		} else {
			left = mid + 1
			res = common.Min(res, notBroken + 1)
		}
	}
	return res
}
