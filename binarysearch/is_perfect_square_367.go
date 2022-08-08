package binarysearch

func IsPerfectSquare(x int) bool {
	if x == 1 {
		return true
	}
	j := x / 2
	i := 0
	for i <= j {
		mid := i + (j - i) / 2
		if mid * mid > x {
			j = mid - 1
		} else if mid * mid < x {
			i = mid + 1
		} else {
			return true
		}
	}
	return false
}
