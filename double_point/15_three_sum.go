package double_point

import (
	"math"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	pre := math.MinInt64
	for i := 0; i < len(nums); i++ {
		// 因为target=0，所以如果nums大于0就不用继续了
		if nums[i] > 0 {
			return res
		}
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left + 1] {
					left++
				}
				for left < right && nums[right] == nums[right - 1] {
					right--
				}
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] > 0{
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// 这种方案解决不了 [0,0,0,0,0] 时重复
func threeSum2(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	pre := math.MinInt64
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > 0 {
			return res
		}
		t := nums[i]
		if t == pre {
			continue
		}
		pre = t
		// 思路是对的，但是取后面两个数的时候有问题
		result := getThreeSumInTwo(nums[i+1:], 0 - t)
		for j := 0; j < len(result); j++ {
			res = append(res, append(result[j], t))
		}
	}
	return res
}

// 这种方案解决不了 [0,0,0,0,0] 时重复
func getThreeSumInTwo(nums []int, target int) [][]int {
	m := make(map[int]int)
	res := make([][]int, 0)
	pre := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			if pre != 0 && nums[i] + pre != target {
				continue
			}
		}
		pre = nums[i]
		remain := target - nums[i]
		if t, exist := m[remain]; exist {
			res = append(res, []int{t, nums[i]})
		}
		m[nums[i]] = nums[i]
	}
	return res
}
