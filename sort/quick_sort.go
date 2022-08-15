package sort

/**
 * 快排, 不稳定排序
 * 最好/平均时间复杂度 O(nlogn)  空间复杂度 O(nlogn)   最差时间复杂度O(n^2)
 * 最好时间复杂度例子：5，3，1，2，4，9，6，8，7，即每次都能均分
 * 最差时间复杂度例子：1，2，3，4，5，6，7，8，9，即基本有序
 */
func quickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	quickSortPart(arr, 0, len(arr) - 1)
}

func quickSortPart(arr []int, start int, end int) {
	// 注意递归终止条件
	if start > end {
		return
	}
	left := start
	right := end
	tmp := arr[start]
	// 注意条件是 < ，而不是 <=
	for left < right {
		// 注意一定是先从后边开始找
		// 注意一定要判断 right > left
		for right > left && arr[right] >= tmp {
			right--
		}
		arr[left] = arr[right]
		for right > left && arr[left] < tmp {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = tmp
	quickSortPart(arr, start, left - 1)
	quickSortPart(arr, left + 1, end)
}
