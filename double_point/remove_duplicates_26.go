package double_point

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	slow := 1
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] != nums[i] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
