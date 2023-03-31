package slide_window

import (
	"algorithm/common"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	drabList := make([]int, 0)
	left, right, valid := 0, 0, 0
	for right < len(nums) {
		n := nums[right]
		right++
		drabList = pushInMaxSlidingWindow(drabList, n)
		valid++
		if valid == k {
			if max, exist := topInMaxSlidingWindow(drabList); exist {
				res = append(res, max)
			}
			m := nums[left]
			// AC这道题很关键的地方，当left缩进的时候，
			// 要把当前最大值给弹出，否则可能在下一轮里无法得到窗口内的最大值
			drabList = popInMaxSlidingWindow(drabList, m) // ！！！
			left++
			valid--
		}
	}
	return res
}

// 单调队列，单调递减
func pushInMaxSlidingWindow(q []int, n int) []int {
	if len(q) == 0 {
		q = append(q, n)
		return q
	}
	// 每次有新值进来，就把队列里比他小的除掉
	for len(q) - 1 >= 0 && q[len(q)-1] < n {
		q = q[:len(q)-1]
	}
	q = append(q, n)
	return q
}

func topInMaxSlidingWindow(q []int) (int, bool) {
	if len(q) == 0 {
		return -1, false
	}
	return q[0], true
}

func popInMaxSlidingWindow(q []int, n int) []int {
	if len(q) == 0 {
		return q
	}
	if q[0] == n {
		q = q[1:]
	}
	return q
}

// 思路没错，但是超时了，
func maxSlidingWindow4(nums []int, k int) []int {
	left := 0
	right := 0
	res := make([]int, 0)
	max := math.MinInt64
	for right < len(nums) {
		n := nums[right]
		if max < n {
			max = n
		}
		if right - left + 1 == k {
			res = append(res, max)
			if nums[left] == max {
				max = math.MinInt64
				for i := left+1; i <= right; i++ {
					max = common.Max(max, nums[i])
				}
			}
			left++
		}
		right++
	}
	return res
}

// 错误的思路，这样每次都变成了取最大值了
func maxSlidingWindow2(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	left, right, valid := 0, 0, 0
	max := math.MinInt64
	res := make([]int, 0)
	for right < len(nums) {
		n := nums[right]
		right++
		valid++
		if max < n {
			max = n
		}
		if valid == k {
			res = append(res, max)
			m := nums[left]
			if max < m {
				max = m
			}
			left++
			valid--
		}
	}
	return res
}

// 理解错题意，以为是计算窗口内值的和
func maxSlidingWindow3(nums []int, k int) []int {
	left, right, valid, sum := 0, 0, 0, 0
	res := make([]int, 0)
	for right < len(nums) {
		n := nums[right]
		right++
		valid++
		sum += n
		if valid == k {
			res = append(res, sum)
			m := nums[left]
			left++
			sum -= m
			valid--
		}
	}
	return res
}
