package math

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n % 3 == 0 {
		n /= 3
	}
	return n == 1
}

// 不使用递归或循环
// 3的n次幂在int类型的最大值是1162261467
// 如果n为3的幂，必然满足 n * 3 ^ k = 1162261467
// 当且仅当 x 为质数时可用，不能快速判断 x 的幂的通用做法
func isPowerOfThree1(n int) bool {
	return n > 0 && 1162261467 % n == 0
}