package binarysearch

func shipWithinDays(weights []int, days int) int {
	maxCarry := 0
	minCarry := 0
	for i := 0; i < len(weights); i++ {
		maxCarry += weights[i]
		if weights[i] > minCarry {
			minCarry = weights[i]
		}
	}
	for minCarry < maxCarry {
		carry := minCarry + (maxCarry - minCarry) / 2
		d := shipWithinDaysConsume(weights, carry)
		if d > days {
			minCarry = carry + 1
		} else {
			maxCarry = carry
		}
	}
	return maxCarry
}

func shipWithinDaysConsume2(weights []int, k int) int {
	d := 0
	sum := 0
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
		if sum > k {
			d++
			sum = weights[i]
		}
		if i + 1 >= len(weights) {
			d++
		}
	}
	return d
}

func shipWithinDaysConsume(weights []int, k int) int {
	d := 0
	for i := 0; i < len(weights); {
		sum := k
		for i < len(weights) {
			if sum < weights[i] {
				break
			} else {
				sum -= weights[i]
			}
			i++
		}
		d++
	}
	return d
}
