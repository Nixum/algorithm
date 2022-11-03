package graph

func findCelebrity(n int) int {
	if n == 1 {
		return 0
	}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		q = append(q, i)
	}
	var cand int
	var other int
	for len(q) >= 2 {
		cand, q = pop(q)
		other, q = pop(q)
		// cand认识other，other不认识cand，说明other可能是名人，但cand一定不是名人
		if know(cand, other) || !know(other, cand) {
			q = append(q, other)
		} else {
			q = append(q, cand)
		}
	}
	// 排除得只剩下一个候选人，最后验证一遍他是不是候选人
	cand, q = pop(q)
	for other := 0; other < n; other++ {
		if other == cand {
			continue
		}
		if know(cand, other) || !know(other, cand) {
			return -1
		}
	}
	return cand
}

// 因为pop是在 len(q) > 2的条件下执行的，所以可以这么写
func pop(q []int) (int, []int) {
	r := q[0]
	q = q[1:]
	return r, q
}

// 预制函数，判断i是否认识j
func know(i, j int) bool {
	return false
}

// 优化版本
func findCelebrity2(n int) int {
	cand := 0
	for other := 1; other < n; other++ {
		if know(cand, other) || !know(other, cand) {
			cand = other
		}
	}
	for other := 0; other < n; other++ {
		if other == cand {
			continue
		}
		if know(cand, other) || !know(other, cand) {
			return -1
		}
	}
	return cand
}