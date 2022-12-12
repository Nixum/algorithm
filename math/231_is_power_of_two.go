package math

// 判断n是否是2的指数次幂的值, 直接 n & (n-1)
// 一个数如果是2的指数，其二进制表示一定只有一个1
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n & (n - 1) == 0
}
