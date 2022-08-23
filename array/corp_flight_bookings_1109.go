package array

// 差分数组 解法
// 比如有原数组：8  2 6  3  1，
//    差分数组：8 -6  4 -3 -2
//   diff[0] = nums[0]; diff[i] = nums[i] - nums[i - 1]
// 所以如果要对nums[i..j]的元素都+3，只需 diff[i] += 3, diff[j+1] -= 3 即可
// 还原：res[i] = diff[0]; res[i] = res[i - 1] + diff[i]
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n + 1)
	for _, booking := range bookings {
		diff[booking[0] - 1] += booking[2]
		diff[booking[1]] -= booking[2]
	}
	res := make([]int, n)
	res[0] = diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i - 1] + diff[i]
	}
	return res
}

// 暴力解法
func corpFlightBookings2(bookings [][]int, n int) []int {
	res := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		b := bookings[i]
		for j := b[0]; j <= b[1]; j++ {
			res[j - 1] += b[2]
		}
	}
	return res
}
