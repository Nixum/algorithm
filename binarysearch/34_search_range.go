package binarysearch

func SearchRange(nums []int, t int) []int {
	return []int{FindInRange(nums, t, false), FindInRange(nums, t, true)}
}

func FindInRange(nums []int, t int, isLeft bool) int {
	i := 0
	j := len(nums) - 1
	res := -1
	for i <= j {
		mid := i + (j - i) / 2
		if nums[mid] > t {
			j = mid - 1
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			if isLeft {
				j = mid - 1
			} else {
				i = mid + 1
			}
			res = mid
		}
	}
	return res
}

// 但是这种时间复杂度是O(n)，只有最好的情况下才是O(logn)
func SearchRange1(nums []int, t int) []int {
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
			res[0]++
			res[1]--
			return res
		}
	}
	return res
}
