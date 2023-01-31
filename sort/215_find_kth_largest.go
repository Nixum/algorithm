package sort

func findKthLargest(nums []int, k int) int {
	for i := len(nums) / 2 - 1; i >= 0; i-- {
		adjustInFindKthLargest(nums, i, len(nums))
	}
	n := k
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustInFindKthLargest(nums, 0, i)
		n--
		if n == 0 {
			return nums[len(nums) - k]
		}
	}
	return 0
}

func adjustInFindKthLargest(nums []int, root int, end int) {
	tmp := nums[root]
	for i := 2 * root + 1; i < end; i = 2 * i + 1{
		if i + 1 < end && nums[i] < nums[i + 1] {
			i++
		}
		if tmp < nums[i] {
			nums[root] = nums[i]
			root = i
		} else {
			break
		}
	}
	nums[root] = tmp
}
