package slide_window

import (
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	drabList := initList()
	left, right, valid := 0, 0, 0
	for right < len(nums) {
		n := nums[right]
		right++
		drabList.push(n)
		valid++
		if valid == k {
			res = append(res, drabList.top())
			m := nums[left]
			// AC这道题很关键的地方，当left缩进的时候，
			// 要把当前最大值给弹出，否则可能在下一轮里无法得到窗口内的最大值
			drabList.pop(m) // ！！！
			left++
			valid--
		}
	}
	return res
}

type drabIncrList struct {
	nums []int
	n int
}

func initList() drabIncrList {
	return drabIncrList{
		nums: make([]int, 0),
		n: 0,
	}
}

func (list *drabIncrList) push(a int) {
	if len(list.nums) == 0 {
		list.nums = append(list.nums, a)
		list.n++
		return
	}
	// 保证单调递增, 每次有新值进来，就把队列里比他小的除掉
	for list.n > 0 && list.nums[list.n - 1] < a {
		list.nums = list.nums[0: list.n - 1]
		list.n--
	}
	list.nums = append(list.nums, a)
	list.n++
}

func (list *drabIncrList) top() int {
	return list.nums[0]
}

func (list *drabIncrList) pop(a int) {
	if list.nums[0] == a {
		if list.n > 1 {
			list.nums = list.nums[1: list.n]
		} else {
			list.nums = list.nums[0:0]
		}
		list.n--
	}
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
