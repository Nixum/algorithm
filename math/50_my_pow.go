package math

// 思路，一个个乘会太慢，那就成倍成倍乘，幂/2，意味着底要平方
// 幂/2 有两种方案，一种是判断n的奇偶，再 x^(n/2) * x^(n/2)，但是这种会超时，虽然指数变小了，但是仍然要算两次
// 另一种是快速幂解法：
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
// 2^5 = 2 ^ (0101) = 按位数直接累乘，n缩小两倍，n末尾为1则额外乘
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

// 这样也会超时, 本质上并没有缩短乘的个数，因为底数还是那么大
func myPow3(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n % 2 == 1 {
		return x * pow(x, n / 2) * pow(x, n / 2)
	}
	return pow(x, n / 2) * pow(x, n / 2)
}
