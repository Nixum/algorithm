package slide_window

func totalFruit(fruits []int) int {
	fruitType := make(map[int]int, 2)
	left := 0
	right := 0
	max := 0
	for right < len(fruits) {
		f := fruits[right]
		fruitType[f]++
		for len(fruitType) > 2 {
			fruitType[fruits[left]]--
			if fruitType[fruits[left]] == 0 {
				delete(fruitType, fruits[left])
			}
			left++
		}
		max = maxN(max, right - left + 1)
		right++
	}
	return max
}

func maxN(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 注意，它是从一颗树开始之后就不能断了，即选择一棵树起的两种类型，能拿到水果的最大数量
// so这个就错了
func totalFruit2(fruits []int) int {
	basket := [2]int{}
	for i := 0; i < len(fruits); i++ {
		if fruits[i] > basket[0] {
			if fruits[i] > basket[1] {
				basket[0] = basket[1]
				basket[1] = fruits[i]
			} else if fruits[i] < basket[1] {
				basket[0] = fruits[i]
			}
		}
	}
	n := 0
	for i := 0; i < len(fruits); i++ {
		if fruits[i] == basket[0] || fruits[i] == basket[1] {
			n++
		}
	}
	return n
}
