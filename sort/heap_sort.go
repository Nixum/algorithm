package sort

func heapSort(arr []int) {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust(arr, 0, i)
	}
}

// 调整成大根堆，根节点大于左右子树的节点
func adjust(arr []int, root int, end int) {
	// 确定tmp的位置
	tmp := arr[root]
	// i 不能小于等于 end，只能小于 end
	for i := 2 * root + 1; i < end; i = 2 * i + 1 {
		if i + 1 < end && arr[i] < arr[i + 1] {
			i = i + 1
		}
		if tmp < arr[i] {
			arr[root] = arr[i]
			root = i
		} else {
			break
		}
	}
	arr[root] = tmp
}
