package week1

// 原理就是让各个位凑成进位，位数变成0和肯定会变小
func makeIntegerBeautiful(n int64, target int) int64 {
	var x int64
	var base int64 = 1 // 从个位开始
	for tran(n) > int64(target) {
		x += (10 - n % 10) * base // 算出要进位的差值
		base *= 10 // 下一位
		n /= 10 // 舍弃低位，因为此时低位是0
		n += 1 // 进位 + 1
	}
	return x
}

func tran(n int64) int64 {
	var res int64 = 0
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res
}

// 暴力解法，超时了
func makeIntegerBeautiful2(n int64, target int) int64 {
	if n == int64(target) {
		return 0
	}
	var res int64 = 0
	for {
		if isBeautiful(n, int64(target)) {
			break
		}

		n++
		res++
	}
	return res
}

func isBeautiful(n int64, target int64) bool {
	var res int64 = 0
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res <= target
}
