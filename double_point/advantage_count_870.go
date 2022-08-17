package double_point

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	// 从大到小排序
	numWithIndexs := make([]numWithIndex, len(nums2))
	for i := 0; i < len(nums2); i++ {
		numWithIndexs[i] = numWithIndex{
			val:   nums2[i],
			index: i,
		}
	}
	sort.Slice(numWithIndexs, func(i, j int) bool {
		return numWithIndexs[i].val > numWithIndexs[j].val
	})
	// 从小到大排序
	sort.Ints(nums1)
	res := make([]int, len(nums1))
	left := 0
	right := len(nums1) - 1
	for i := 0; i < len(numWithIndexs); i++ {
		// 田忌赛马
		// 如果nums1大的，就作为结果，如果还小，那就拿最小的作为结果
		if nums1[right] > numWithIndexs[i].val {
			res[numWithIndexs[i].index] = nums1[right]
			right--
		} else {
			res[numWithIndexs[i].index] = nums1[left]
			left++
		}
	}
	return res
}

type numWithIndex struct {
	val int
	index int
}
