package binarysearch

func SearchRange(nums []int, t int) []int {
	res := []int{-1, -1}
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := i + (j - i) / 2
		if nums[mid] > t {
			j = mid - 1
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			res[0] = mid
			res[1] = mid
			for nums[res[0]] == t {
				res[0]--
			}
			for nums[res[1]] == t {
				res[1]++
			}
			return res
		}
	}
	return res
}
