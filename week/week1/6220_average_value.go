package week1

func averageValue(nums []int) int {
	res := 0
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] % 3 == 0 && nums[i] % 2 == 0 {
			res += nums[i]
			n++
		}
	}
	if n != 0 {
		return res / n
	} else {
		return 0
	}
}
