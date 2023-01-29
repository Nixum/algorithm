package math

// 只出现一次的数字
// 一个数和它本身做异或运算结果为0，即a^a=0
// 一个数和0做异或运算结果为它本身，即a^0=a
func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
