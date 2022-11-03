package backtracking

// 划分成k个相等的子集，每个子集的和相等
func canPartitionKSubsets(nums []int, k int) bool {
	if k <= 0 {
		return false
	}
	if len(nums) < k {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum % k != 0 {
		return false
	}
	target := sum / k
	isUsed := make([]bool, len(nums))
	return backTrackingCanPartitionKSubsets(nums, k, 0, 0, target, isUsed)
}

func backTrackingCanPartitionKSubsets(nums []int, k int, start int, bucketSum int, target int, isUsed []bool) bool {
	if k == 0 {
		return true
	}
	if bucketSum == target {
		return backTrackingCanPartitionKSubsets(nums, k - 1, 0, 0, target, isUsed)
		// 这样写会导致进不了下一个桶的循环
		//k--
		//bucketSum = 0
	}
	// 寻找每一个桶内存放的数字
	for i := start; i < len(nums); i++ {
		t := bucketSum + nums[i]
		if t > target || isUsed[i] {
			continue
		}
		bucketSum += nums[i]
		isUsed[i] = true
		res := backTrackingCanPartitionKSubsets(nums, k, i + 1,bucketSum, target, isUsed)
		if res {
			return true
		}
		isUsed[i] = false
		bucketSum -= nums[i]
	}
	return false
}