package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1,-1}, 1))
	fmt.Println(maxSlidingWindow([]int{7,2,4}, 2))
	fmt.Println(maxSlidingWindow([]int{1,3,1,2,0,5}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := 0
	ascArr := make([]int, 0)
	res := make([]int, 0)
	for right < len(nums) {
		n := nums[right]
		ascArr = push(ascArr, n)
		if right - left + 1 == k {
			if max, exist := top(ascArr); exist {
				res = append(res, max)
			}
			ascArr = pop(ascArr, nums[left])
			left++
		}
		right++
	}
	return res
}

// 维护单调递减队列
func push(q []int, n int) []int {
	if len(q) == 0 {
		q = append(q, n)
		return q
	}
	for len(q) - 1 >= 0 && q[len(q)-1] < n {
		q = q[:len(q)-1]
	}
	q = append(q, n)
	return q
}

func top(q []int) (int, bool) {
	if len(q) == 0 {
		return -1, false
	}
	return q[0], true
}

func pop(q []int, n int) []int {
	if len(q) == 0 {
		return q
	}
	if q[0] == n {
		q = q[1:]
	}
	return q
}
