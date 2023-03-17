package array

// 找到数组的下一个排列
// 123456 、123465 、123546 ... 654321
// 期望：
// 1. 下一个数比当前数尽可能大 - 找降序分界点，第一个后面的大数跟前面的小数交换
// 2. 增加的幅度尽可能小 - 升序，使其尽可能大，才能越接近当前数
// 思路：
// 1. 从后向前，找到第一个 升序对 nums[i] < nums[j]，如果找不到说明符合，此时直接进行排序即可
// 此时 nums[j-end] 都是降序的
// 2. 从 nums[j-end] 找到第一个比 nums[i] 大的数 nums[k]
// 3. 交换，此时 nums[j-end] 也是降序
// 4. 对 nums[j-end] 的数组进行升序
func nextPermutation(nums []int)  {
	start := 0
	// 找第一个升序
	for i := len(nums)-1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			start = i
			break
		}
	}
	if start != 0 {
		for i := len(nums)-1; i >= start; i-- {
			// 找到大数
			if nums[i] > nums[start-1] {
				nums[start-1], nums[i] = nums[i], nums[start-1]
				break
			}
		}
	}
	// 升序
	end := len(nums)-1
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return
}
