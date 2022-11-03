package binarysearch

func SearchInsert(nums []int, t int) int {
	i := 0
	j := len(nums) - 1
	mid := -1
	for i <= j {
		mid = i + (j - i) / 2
		if nums[mid] > t {
			j = mid - 1
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			return mid
		}
	}
	if nums[mid] < t {
		return mid + 1
	}
	return mid - 1
}