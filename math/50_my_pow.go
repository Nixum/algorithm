package math

// 思路，一个个乘会太慢，那就成倍成倍乘，幂/2，意味着底要平方
// 快速幂解法：
// 如果幂是偶数，以 2^4为例子，2^4 = (2^2)^2 = ((2^1)^2)^2
// 如果幂是奇数，则把奇数进行拆分，拆一个出来乘即可
// 时间复杂度，空间复杂度 都是 O(logn)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n % 2 == 0 {
		return myPow(x * x, n / 2)
	}
	return x * myPow(x, n - 1)
	// 或者写成下面这样
	//return x * myPow(x * x, n / 2)
}

// 二进制法，把n转成二进制数
// 2^5 = 2 ^ (0101) = 平时直接累乘，碰到1则额外乘
// 时间复杂度是 O(logn), 空间复杂度1
func myPow2(x float64, n int) float64 {
	isNative := false
	if n < 0 {
		isNative = true
		n = -n
	}
	var res float64 = 1
	for n > 0 {
		if n % 2 == 1 {
			res *= x
		}
		x *= x
		n = n / 2
	}
	if isNative {
		return 1 / res
	}
	return res
}

// 这样会超时，因为需要一个个乘
func myPow1(x float64, n int) float64 {
	isNative := false
	if n < 0 {
		isNative = true
		n = -n
	}
	res := float64(1)
	for i := 0; i < n; i++ {
		res *= x
	}
	if isNative {
		return 1/res
	}
	return res
}
