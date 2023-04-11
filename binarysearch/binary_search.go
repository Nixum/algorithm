package binarysearch

func BinarySearch(nums []int, t int) int {
	i := 0
	j := len(nums) - 1
	for i <= j { // 注意 j = len(nums) 时无需比较 相等
		mid := i + (j - i) / 2 // 防止溢出
		if nums[mid] > t {
			j = mid - 1  // 注意 j = len(nums) 时无需 - 1
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 要注意的一个点是，因为 除法是向下取整，当i要等于mid时，
// 可能会导致死循环，所以 要不就使用 mid = (l+r+1)/2, 要不就 j=mid，i=mid+1
func BinarySearch2(nums []int, t int) int{
	i := 0
	j := len(nums)
	for i < j {
		mid := i + (j - i) / 2
		if nums[mid] > t {
			j = mid
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
