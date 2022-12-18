package math

// 统计所有小于非负整数n的质数的数量
func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	// 统计非质数的个数
	for i := 2; i < n; i++ {
		// 如果是质数
		if !isNotPrime[i] {
			// 计算该质数的倍数的值为非质数
			for j := 2; j * i < n; j++ {
				isNotPrime[j * i] = true
			}
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			res++
		}
	}
	return res
}
