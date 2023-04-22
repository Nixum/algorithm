package array

import "sort"

// 给你两个整数数组 nums1 和 nums2，请你以数组形式返回两数组的交集。
// 返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）
// 就是要求是交集，且次数一致
func intersect(nums1 []int, nums2 []int) []int {
	arr1 := make([]int, 1001)
	for i := 0; i < len(nums1); i++ {
		arr1[nums1[i]]++
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if arr1[nums2[i]] > 0 {
			res = append(res, nums2[i])
			arr1[nums2[i]]--
		}
	}
	return res
}

// 进阶，如果nums1和nums2已排完序
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	left := 0
	right := 0
	res := make([]int, 0)
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] == nums2[right] {
			res = append(res, nums1[left])
			left++
			right++
		} else if nums1[left] > nums2[right] {
			right++
		} else {
			left++
		}
	}
	return res
}