package binarysearch

// 求x的算术平方根，由于结果是整形，所以返回去掉小数的结果
func MySqrt(x int) int {
	if x == 1 {
		return 1
	}
	for i := 1; i <= x / 2; i++ {
		if i * i <= x && (i + 1) * (i + 1) > x {
			return i
		}
	}
	return 0
}

// 思路：二分查找，返回小于x的最大值或者等于x的值
func MySqrt2(x int) int {
	if x == 1 {
		return 1
	}
	j := x / 2
	i := 0
	res := -1
	for i <= j {
		mid := i + (j - i) / 2
		if mid * mid > x {
			j = mid - 1
		} else if mid * mid < x {
			res = mid
			i = mid + 1
		} else {
			return mid
		}
	}
	return res
}
