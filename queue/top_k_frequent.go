package queue

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	m := make(map[int]PairTopKFrequent)
	// 统计出现的频率
	for i := 0; i < len(nums); i++ {
		if obj, exist := m[nums[i]]; exist {
			obj.freq++
			m[nums[i]] = obj
		} else {
			m[nums[i]] = PairTopKFrequent{
				num:  nums[i],
				freq: 1,
			}
		}
	}
	pairs := make([]PairTopKFrequent, 0)
	for _, obj := range m {
		pairs = append(pairs, obj)
	}
	res := make([]int, 0)
	for i := len(pairs) / 2 - 1; i >= 0; i-- {
		adjustTopKFrequent(pairs, i, len(pairs) - 1)
	}
	for i := len(pairs) - 1; i >= 0; i-- {
		if k == 0 {
			break
		}
		pairs[0], pairs[i] = pairs[i], pairs[0]
		adjustTopKFrequent(pairs, 0, i - 1)
		res = append(res, pairs[i].num)
		k--
	}
	return res
}

type PairTopKFrequent struct {
	num int
	freq int
}

// 大根堆，单调递减排序
func adjustTopKFrequent(nums []PairTopKFrequent, rootIndex int, end int) {
	tmp := nums[rootIndex]
	for i := 2 * rootIndex + 1; i <= end; i = 2*i+1 {
		if i + 1 <= end && nums[i].freq < nums[i + 1].freq {
			i++
		}
		if nums[i].freq > tmp.freq {
			nums[rootIndex] = nums[i]
			rootIndex = i
		} else {
			break
		}
	}
	nums[rootIndex] = tmp
}