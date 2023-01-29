package math

// 返回其二进制表达式中数字位数为 '1' 的个数
// n & (n-1)：作用是消除数字n的二进制数中的最后一个1
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num &= num-1
		res++
	}
	return res
}
