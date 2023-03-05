package sort

// 时间复杂度是O(n), 空间复杂度是O(1), 排列只有 0，1，2的数组
func sortColors(nums []int)  {
	r := 0
	w := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			r++
		case 1:
			w++
		case 2:
			b++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < r {
			nums[i] = 0
		} else if i < (r + w) {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
