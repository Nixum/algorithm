package math

import (
	"algorithm/common"
	"math"
)

// 要求时间复杂度是O(log(n+m))
// 思路：
// 求中位数，说明要取的个数是 (len(nums1)+len(nums2)+1)/2
// 因为nums1和nums2都是升序，那只要找到nums1和nums2的临界值，就是中位数的位置了
// 比如对 nums1 做二分查找，判断 nums1[mid] 和 nums2[(len(nums1)+len(nums2)+1)/2-mid] 的大小
// 直到找到临界点，此时临界点左边是max(nums1[mid-1], nums2[mid-1])
// 临界点右边是 min(nums1[mid], nums2[mid])
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (len1+len2+1)/2
	left := 0
	right := len1
	for left < right {
		mid := left + (right - left) / 2
		if nums1[mid] < nums2[k-mid-1] {
			// nums1的元素不够多
			left = mid + 1
		} else {
			// nums1的元素太多了
			right = mid
		}
	}
	m1 := left
	m2 := k-left
	c1 := math.MinInt64
	if m1 > 0 {
		c1 = nums1[m1-1]
	}
	if m2 > 0 {
		c1 = common.Max(c1, nums2[m2-1])
	}

	if (len1 + len2)%2 == 1 {
		return float64(c1)
	}

	c2 := math.MaxInt64
	if m1 < len1 {
		c2 = nums1[m1]
	}
	if m2 < len2 {
		c2 = common.Min(c2, nums2[m2])
	}
	return float64(c1 + c2) / 2
}

// 时间复杂度是O((n+m)/2)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	sumLen := len(nums1) + len(nums2)
	res1 := 0
	res2 := 0
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		res1 = res2
		if nums1[i] <= nums2[j] {
			res2 = nums1[i]
			i++
		} else {
			res2 = nums2[j]
			j++
		}
		if i + j > sumLen/2 {
			break
		}
	}
	for i < len(nums1) && i + j <= sumLen/2 {
		res1 = res2
		res2 = nums1[i]
		i++
	}
	for j < len(nums2) && i + j <= sumLen/2 {
		res1 = res2
		res2 = nums2[j]
		j++
	}
	if sumLen % 2 == 0 {
		return float64(res1 + res2) / 2
	} else {
		return float64(res2)
	}
}

// 时间复杂度是O(n+m)，空间复杂度也是
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	sumLen := len(nums1) + len(nums2)
	nums := make([]int, sumLen)
	index1 := 0
	index2 := 0
	index := 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			nums[index] = nums1[index1]
			index1++
		} else if nums1[index1] > nums2[index2] {
			nums[index] = nums2[index2]
			index2++
		} else {
			nums[index] = nums1[index1]
			index++
			nums[index] = nums2[index2]
			index1++
			index2++
		}
		index++
	}
	for index1 < len(nums1) {
		nums[index] = nums1[index1]
		index++
		index1++
	}
	for index2 < len(nums2) {
		nums[index] = nums2[index2]
		index++
		index2++
	}
	if sumLen % 2 == 0 {
		return float64(nums[sumLen/2] + nums[sumLen/2 - 1]) / 2
	} else {
		return float64(nums[sumLen/2])
	}
}
