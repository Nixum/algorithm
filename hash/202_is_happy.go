package hash

func isHappy(n int) bool {
	set := make(map[int]int)
	for true {
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		if _, exist := set[sum]; exist {
			break
		}
		set[sum] = 1
		n = sum
	}
	return false
}

func getSum(n int) int {
	sum := 0
	for n != 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	return sum
}
