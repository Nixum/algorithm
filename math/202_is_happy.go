package math

// 快慢指针解决，用map去重可能会导致map特别大
// 因为无论是里面含有计算循环，还是最终等于1，都会导致fast == slow，此时可以退出外层循环
// 然后再判断最终结果是否是1来判断是否是快乐数
func isHappy(n int) bool {
	slow := n
	fast := calInIsHappy(n)
	for slow != fast {
		slow = calInIsHappy(slow)
		fast = calInIsHappy(calInIsHappy(fast))
	}
	return slow == 1
}

func isHappy1(n int) bool {
	dup := make(map[int]int)
	for n != 1 {
		n = calInIsHappy(n)
		if _, exist := dup[n]; exist {
			return false
		}
		dup[n] = 1
	}
	return true
}

func calInIsHappy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
