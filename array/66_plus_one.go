package array

// 优雅一点的写法
func plusOne(digits []int) []int {
	for i := len(digits)-1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	newDigits := make([]int, 1, len(digits)+1)
	newDigits[0] = 1
	newDigits = append(newDigits, digits...)
	return newDigits
}

func plusOne1(digits []int) []int {
	next := 0
	for i := len(digits)-1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i]++
		}
		digits[i] += next
		if digits[i] == 10 {
			next = 1
			digits[i] = 0
		} else {
			next = 0
		}
		if next == 0 {
			break
		}
	}
	if next != 0 {
		newDigits := make([]int, 1, len(digits)+1)
		newDigits[0] = 1
		newDigits = append(newDigits, digits...)
		return newDigits
	}
	return digits
}