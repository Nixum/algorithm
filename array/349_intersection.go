package array

// 求两数组交集
func intersection349(nums1 []int, nums2 []int) []int {
	nums := make([]int, 1001)
	for i := 0; i < len(nums1); i++ {
		nums[nums1[i]]++
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if nums[nums2[i]] > 0 && nums[nums2[i]] != -1 {
			res = append(res, nums2[i])
			nums[nums2[i]] = -1
		}
	}
	return res
}

// 进阶，两个数组是有序数组A和B，且包含重复元素，判断B是否是A的子集
// 双指针解决
func intersection1(nums1 []int, nums2 []int) bool {
	left := 0
	right := 0
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] == nums2[right] {
			left++
			right++
		} else if nums1[left] > nums2[right] {
			return false
		} else {
			left++
		}
	}
	return len(nums2) == right
}