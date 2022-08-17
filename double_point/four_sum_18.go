package double_point

import (
	"math"
	"sort"
)

func fourSum(nums []int, t int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	pre := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		pre2 := math.MinInt64
		for j := i + 1; j < len(nums); j++ {
			if pre2 == nums[j] {
				continue
			}
			pre2 = nums[j]
			left := j + 1
			right := len(nums) - 1
			for left < right {
				if nums[i] + nums[j] + nums[left] + nums[right] == t {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[i] + nums[j] + nums[left] + nums[right] > t {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}
