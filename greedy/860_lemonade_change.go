package greedy

func lemonadeChange(bills []int) bool {
	coins := make([]int, 2)
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			coins[0]++
		case 10:
			coins[0]--
			if coins[0] < 0 {
				return false
			}
			coins[1]++
		case 20:
			if coins[1] - 1 >= 0 && coins[0] - 1 >= 0 {
				coins[1]--
				coins[0]--
				continue
			}
			if coins[1] - 1 < 0 && coins[0] - 3 >= 0 {
				coins[0] -= 3
				continue
			}
			return false
		}
	}
	return true
}