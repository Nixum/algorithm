package hash

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := hash[nums[i]]; exist {
			return true
		}
		hash[nums[i]]++
	}
	return false
}
