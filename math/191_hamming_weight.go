package math

// n & (n-1)：作用是消除数字n的二进制数中的最后一个1
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num &= num-1
		res++
	}
	return res
}
