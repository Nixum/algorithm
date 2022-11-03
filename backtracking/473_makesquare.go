package backtracking

func makesquare(matchsticks []int) bool {
	sum := 0
	for i := 0; i < len(matchsticks); i++ {
		sum += matchsticks[i]
	}
	if sum % 4 != 0 {
		return false
	}
	resInMakeSquare = false
	isUsedInMakeSquare := make([]bool, len(matchsticks))
	target := sum / 4
	backTrackingInMakeSquare(matchsticks, 0, 4, 0, target, isUsedInMakeSquare)
	return resInMakeSquare
}

var resInMakeSquare bool
func backTrackingInMakeSquare(matchsticks []int, start, k, edgeSum, target int, used []bool) {
	if k == 0 {
		resInMakeSquare = true
		return
	}
	if resInMakeSquare {
		return
	}
	if edgeSum == target {
		backTrackingInMakeSquare(matchsticks, 0, k - 1, 0, target, used)
		return
	}
	for i := start; i < len(matchsticks); i++ {
		if used[i] {
			continue
		}
		if edgeSum + matchsticks[i] > target {
			continue
		}
		edgeSum += matchsticks[i]
		used[i] = true
		backTrackingInMakeSquare(matchsticks, i + 1, k, edgeSum, target, used)
		if resInMakeSquare {
			return
		}
		used[i] = false
		edgeSum -= matchsticks[i]
	}
}
