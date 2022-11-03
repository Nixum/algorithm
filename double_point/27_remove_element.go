package double_point

func removeElement(nums []int, val int) int {
	replace := 0
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			num++
		} else {
			nums[replace] = nums[i]
			replace++
		}
	}
	return len(nums) - num
}
