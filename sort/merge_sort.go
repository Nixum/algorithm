package sort

// 归并排序，时间复杂度：O(nlogn)，空间复杂度：O(n)，稳定排序
func mergeSort(arr []int) {
	sort(arr, 0, len(arr) - 1)
}

func sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end - start) / 2
	sort(arr, start, mid)
	sort(arr, mid + 1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, 0)
	s := start
	m := mid + 1
	for s <= mid && m <= end {
		if arr[s] < arr[m] {
			tmp = append(tmp, arr[s])
			s++
		} else {
			tmp = append(tmp, arr[m])
			m++
		}
	}
	for s <= mid {
		tmp = append(tmp, arr[s])
		s++
	}
	for m <=end {
		tmp = append(tmp, arr[m])
		m++
	}
	for i := 0; i < len(tmp); i++ {
		arr[start] = tmp[i]
		start++
	}
}