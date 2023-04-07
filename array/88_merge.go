package array

// 末尾追加
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m-1
	j := n-1
	le := m + n -1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[le] = nums2[j]
			j--
		} else {
			nums1[le] = nums1[i]
			i--
		}
		le--
	}
	// 如果nums2有剩余，copy到nums1即可
	for s := 0; s < j+1; s++ {
		nums1[s] = nums2[s]
	}
	// 如果nums1有剩余，因为nums1已经在nums1里了，不用copy
}
