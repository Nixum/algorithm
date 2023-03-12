package math


// 给一个整数 n ，对于 0 <= i <= n 中的每个 i ，
// 计算其二进制表示中 1 的个数

// i 有两种情况，
// 偶数，比如2：10，4：100，8：1000，二进制末位是0，所以二进制中1的个数没有变，右移一位之后等于i/2，1的个数没有变
// 奇数，比如1：01，3：011，5：101，二进制末位是1，保底就有一个1，减去末位的1之后，就变成偶数，偶数算法就跟上面的一样了
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			res[i] = res[(i-1)/2] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}


func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = count(i)
	}
	return res
}

func count(n int) int {
	res := 0
	for n != 0 {
		n &= n-1
		res++
	}
	return res
}
