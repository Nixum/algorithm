package double_point

// 时间复杂度是O(n), 空间复杂度是O(1), 排列只有 0，1，2的数组
// 思路是先把所有数都置为2，然后置为1，再置为0，置为0和1用双指针
func sortColors(nums []int)  {
	i0, i1 := 0, 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		nums[i] = 2
		// 必须先比较2再比较1
		if n < 2 {
			nums[i1] = 1
			i1++
		}
		if n < 1 {
			nums[i0] = 0
			i0++
		}
	}
}

// 三指针，思路是利用三指针对数组分成三个区间
func sortColors2(nums []int)  {
	s := 0
	i := 0
	e := len(nums)-1
	for i <= e {
		switch nums[i] {
		case 0:
			nums[s], nums[i] = nums[i], nums[s]
			s++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[e] = nums[e], nums[i]
			e--
		}
	}
}