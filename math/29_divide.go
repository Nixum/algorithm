package math

import "math"

// 实现除法，不允许使用乘法、除法、取模
// 思路：倍增乘法，利用对于 x 除以 y，结果必然落在范围 [0, x] 的规律进行二分
func divide(dividend int, divisor int) int {
	x := dividend
	y := divisor
	isNegate := false
	if (x > 0 && y < 0) || (x < 0 && y > 0) {
		isNegate = true
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	l := 0
	r := x
	// 二分查找 [0, x]的值，判断哪个数 乘以 除数 会等于 被除数
	for l < r {
		mid := l + (r - l)/2 + 1
		tmp := multi(mid, y)
		if tmp <= x {
			l = mid
		} else {
			r = mid-1
		}
	}
	ans := l
	if isNegate {
		ans = -l
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt64
	}
	return ans
}

// 乘法实现
func multi(a, b int) int {
	ans := 0
	for b > 0 {
		if b & 1 == 1 {
			ans += a
		}
		b >>= 1
		a += a
	}
	return ans
}

// 倍乘的另一种写法
// 假定被除数为 20，除数为 3，使用倍乘法的过程如下
// 计算 3的2^x的最大值为 3 * 2^2 = 12, 20-12=8, 8作为新的被除数
// 计算 3的2^x的最大值为 3 * 2^1 = 6, 8-6=2, 2作为新的被除数
// 2 < 3, 退出计算过程，结果为 2^2 + 2^1 = 6
func divide2(dividend int, divisor int) int {
	x := dividend
	y := divisor
	// 负数的最小值的绝对值比正数的最大值大 1，
	// 所以，直接取反是会溢出，这种情况特殊处理
	if x == math.MinInt32 && y == -1 {
		return math.MaxInt32
	}
	isNotNegate := true
	if (x > 0 && y > 0) || (x < 0 && y < 0) {
		isNotNegate = false
	}
	// 全部转换成负数，防止溢出
	if x > 0 {
		x = -x
	}
	if y > 0 {
		y = -y
	}
	ans := 0
	for x <= y {
		tmp := y
		count := 1
		for tmp >= x-tmp {
			tmp += tmp
			count += count
		}
		// 被除数减去除数的 2^x 倍数做为新的被除数
		x -= tmp
		ans += count
	}
	if isNotNegate {
		return ans
	}
	return -ans
}

// 思路没错，但是超时了，一个个减太费时了
func divide1(dividend int, divisor int) int {
	isNegate := false
	if dividend < 0 && divisor > 0 {
		dividend = -dividend
		isNegate = true
	}
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		isNegate = true
	}
	if dividend < 0 && divisor < 0 {
		divisor = -divisor
		dividend = -dividend
	}
	res := 0
	if divisor == 1 {
		res = dividend
	} else {
		for dividend > 0 {
			if dividend < divisor {
				break
			}
			dividend -= divisor
			res++
		}
	}
	if !isNegate && res > math.MaxInt32 {
		return math.MaxInt32
	}
	if isNegate && -res < math.MinInt32 {
		return math.MinInt64
	}
	if isNegate {
		return -res
	}
	return res
}
