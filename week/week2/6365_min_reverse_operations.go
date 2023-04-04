package week2

// 翻转之后的下标范围, 1 可以是从 i-k+1 到 i+k-1 的一个公差为2的等差数列
// 范围是L，R，L=0, R=n-1，利用对称性，比如 i 和 ii 是对称，他们的下标加起来一定是一个常数,
// 左边界，ii+i = k-1, 右边界，ii = (n-1)+(n-k)-i
// 那范围就是 左边界：max(i-k+1, k-i-1)， 右边界：min(i+k-1, (n-1)+(n-k)-i)
// 配合平衡树解决
func minReverseOperations(n int, p int, banned []int, k int) []int {

	return []int{}
}
