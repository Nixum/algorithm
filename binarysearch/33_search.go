package binarysearch

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		}
		// 无法根据target和nums[mid]来判断target在哪个区间
		// 所以先判断mid两边是什么区间
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
