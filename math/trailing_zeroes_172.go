package math

// 判断n的阶乘尾部有多少个0
// 尾部为0，说明因子中有2和5，而2能被偶数整除
// 问题就变成了n的阶乘能分解出多少个因子5
func trailingZeroes(n int) int {
	res := 0
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	return res
}
