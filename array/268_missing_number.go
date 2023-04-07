package array

// 要求时间复杂度为O(n)，空间复杂度为O(1)
func missingNumber(nums []int) int {
	nums = append(nums, -1)
	for i := 0; i < len(nums); i++ {
		for nums[i] != i && nums[i] != -1 {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			return i
		}
	}
	return 0
}

// 还有其他方式，比如求 1+到n的总和 - 数组总和
// 对1到n进行异或得到ans，然后用ans对各个nums[i]进行异或