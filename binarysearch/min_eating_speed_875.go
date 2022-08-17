package binarysearch

import "math"

// 吃香蕉的速度k，只能大于 所有香蕉总数/总小时，小于香蕉最大数
// 所以只需要通过二分法在这个范围内找到最接近h的值，就是吃香蕉最小速度
func minEatingSpeed(piles []int, h int) int {
	sum := 0
	max := -1
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
		if piles[i] > max {
			max = piles[i]
		}
	}
	min := sum / h
	for min < max {
		curK := min + (max - min) / 2
		kh := hourRequired(piles, curK)
		// 如果用curK的速度消耗完这堆香蕉需要的h大于目标值
		// 说明消耗速度太慢，需要加快速度
		if kh > h {
			min = curK + 1
		} else {
			max = curK
		}
	}
	return max
}

func hourRequired(piles []int, k int) int {
	h := 0
	if k == 0 {
		return math.MaxInt64
	}
	for i := 0; i < len(piles); i++ {
		if piles[i] % k != 0 {
			h++
		}
		h += piles[i] / k
	}
	return h
}
