package greedy

func wiggleMaxLength(nums []int) int {
	preDiff := 0
	curDiff := 0
	res := 1
	for i := 0; i < len(nums) - 1; i++ {
		curDiff = nums[i + 1] - nums[i]
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}

// 错误的思路
func wiggleMaxLength2(nums []int) int {
	res := 1
	isNegative := false
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	if nums[1] - nums[0] < 0 {
		isNegative = true
	}
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i + 1] {
			continue
		}
		if isNegative {
			if nums[i + 1] - nums[i] > 0 {
				isNegative = false
				res++
			}
		} else {
			if nums[i + 1] - nums[i] < 0 {
				isNegative = true
				res++
			}
		}
	}
	return res
}
