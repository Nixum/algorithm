package double_point

import "fmt"

// water[i] = min(max(height[0-i]), max(height[i-end])) - height[i]
func trap(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	lMax := 0
	rMax := 0
	for left <= right {
		if lMax < height[left] {
			lMax = height[left]
		}
		if rMax < height[right] {
			rMax = height[right]
		}
		// 关键点，比如有2, 0, 7(或者1), 5,
		// 因为是左右两边取最大值再取最小值，所以中间的值是不影响的
		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}

	}
	return res
}

// 比较好理解的方案
func trap2(height []int) int {
	hLen := len(height)
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))
	lMax[0] = height[0]
	lhMax := height[0]
	rMax[hLen - 1] = height[hLen - 1]
	rhMax := height[hLen - 1]
	res := 0
	// 简化版本实际上可以把两个数组去掉，只留下最值比较，通过双指针实现
	for i := 1; i < hLen; i++ {
		if lhMax < height[i - 1] {
			lhMax = height[i - 1]
		}
		lMax[i] = lhMax
	}
	for i := hLen - 2; i >= 0; i-- {
		if rhMax < height[i + 1] {
			rhMax = height[i + 1]
		}
		rMax[i] = rhMax
	}
	fmt.Println(lMax)
	fmt.Println(rMax)
	fmt.Println(height)
	for i := 0; i < hLen; i++ {
		res += max(min(lMax[i], rMax[i]) - height[i], 0 )
		fmt.Print(min(lMax[i], rMax[i]) - height[i], " ")
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
