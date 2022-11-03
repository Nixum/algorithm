package double_point

func moveZeroes(nums []int)  {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
