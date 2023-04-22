package binarysearch

// 因为存在重复元素，所以没办法像33题那样一下子能区分前面有序还是后面有序
// 所以针对 10111 或 11101 这种情况，当nums[left] == nums[mid]时，left++即可，先排除掉一个干扰项
// 其他情况就跟 33题 一样了
func search81(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[left] {
			if target < nums[mid]  && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return false
}

// 另一种解法，直接遍历即可，因为存在重复元素，极端情况下，使用上面那种算法时间复杂度也是O(n)