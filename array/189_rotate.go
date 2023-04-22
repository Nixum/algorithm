package array

// 翻转法
func rotate13(nums []int, k int)  {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	i := 0
	j := len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// 要求空间复杂度为O(1)
func rotate1(nums []int, k int)  {
	k %= len(nums)
	count := gcd(k, len(nums))
	for start := 0; start < count; start++ {
		pre, curIndex := nums[start], start
		for {
			nextIndex := (curIndex + k) % len(nums)
			nums[nextIndex], pre = pre, nums[nextIndex]
			curIndex = nextIndex
			if curIndex == start {
				break
			}
		}
	}
}

// 求最大公约数
// 12 和 16，最大公约数是4
func gcd(a, b int) int {
	// 辗转相除法
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 没有用最大公约数，也能ac，就是会多循环几次
func rotate12(nums []int, k int)  {
	k %= len(nums)
	cnt := 0
	for start := 0; start < len(nums); start++ {
		pre, curIndex := nums[start], start
		for {
			nextIndex := (curIndex + k) % len(nums)
			nums[nextIndex], pre = pre, nums[nextIndex]
			curIndex = nextIndex
			cnt++
			if curIndex == start {
				break
			}
		}
		if cnt == len(nums) {
			break
		}
	}
}

// 思路没错，但是没有解决轮转多一次的问题
func rotate111(nums []int, k int)  {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		nextIndex := -1
		index := i
		for index + k < len(nums) && nextIndex != index {
			nextIndex = (index + k)%len(nums)
			tmp := nums[nextIndex]
			nums[nextIndex] = nums[index]
			nums[index] = tmp
			index += k
			cnt++
		}
		if cnt == len(nums) {
			break
		}
	}
}