package double_point

import (
	"math"
	"sort"
)

func twoSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	pre := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if pre == nums[i] {
			continue
		}
		pre = nums[i]
		right := len(nums) - 1
		for i < right {
			if nums[i] + nums[right] == target {
				res = append(res, []int{nums[i], nums[right]})
				for right > i && nums[right] == nums[right - 1] {
					right--
				}
			}
			right--
		}
	}
	return res
}
