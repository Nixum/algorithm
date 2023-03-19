package math

// 寻找数组中重复的数
// 要求不修改原数组，最好能实现时间复杂度为O(n)，空间复杂度为O(1)
// 思路：把这种场景看成带有环的链表
// 比如有数组：1 3 2 2 4
// 映射成 0-1，1-3，2-2，3-2，4-2，2-2就是一个环，4-2走不到就不管它，因为之前已经出现环了
// 那就可以用快慢指针来解决
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

// 时间复杂度是 O(nlogn)的解法
// 思路，对1-n做二分查找 猜数字，配合抽屉原理进行判断
// 抽屉原理，10个苹果放入9个抽屉，必定有一个抽屉有两个
func findDuplicate2(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left < right {
		mid := left + (right - left)/2

		count := 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		// 1 2 3 3 4
		// mid = 3, count = 4, 重复的数出现在左区间
		// 1 2 2 3 4
		// mid = 2, count = 3, 重复的数出现在左区间
		// 1 2 3 3 4
		// mid = 3, count = 3, 重复的数出现在右区间
		// 1 2 3 4 4
		// mid = 3, count = 3, 重复的数出现在右区间
		if count <= mid {
			left = mid+1
		} else {
			right = mid
		}
	}
	return left
}
