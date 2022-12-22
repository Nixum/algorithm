package math

// 求x的算术平方根，由于结果是整形，所以返回去掉小数的结果
// 思路：二分查找，返回小于x的最大值或者等于x的值
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left := 1
	right := x/2
	for left < right {
		mid := left + (right - left + 1) / 2
		midRes := mid * mid
		if midRes > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
