package array

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
	left := 0
	right := 0
	res := make([]int, 0)
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] == nums2[right] {
			res = append(res, nums1[left])
		} else if nums1[left] > nums2[right] {
			right++
		} else {
			left++
		}
	}
	return res
}