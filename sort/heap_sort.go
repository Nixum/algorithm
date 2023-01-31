package sort

/**
 * 堆排序
 * 两个步骤：1.创建初始堆    2.头节点与末结点交换，调整满足大根堆定义
 * 如何创建初始堆：
 *      从第一个非叶子结点开始，比较该结点的左右孩子，大的结点与其父结点交换，
 *      沿着交换的叶子结点的左右孩子继续判断是否满足大根堆定义，调整；
 *      第一个非叶子结点的下标减1，重复调整操作，直至调整满足大根堆定义
 * 如何调整：
 *      从数组最后一个元素开始，之后每次减1，与根结点交换
 *      根结点为最大值，将其与最后一个结点交换，此时最后一个结点排最后了，
 *      然后从根结点开始，调整至满足大根堆定义，此时根结点为次最大值
 *      根结点与倒数第二个结点交换，调整，如此反复，即完成排序
 *
 * 最差/平均时间复杂度 O(nlogn)
 * 最好时间复杂度O(nlogn)
 * 空间复杂度 O(1)
 */
func heapSort(arr []int) {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		// 排除掉i之后的不调整
		adjust(arr, 0, i)
	}
}

// 调整成大根堆，根节点大于左右子树的节点
func adjust(arr []int, root int, end int) {
	// 确定tmp的位置
	tmp := arr[root]
	// i 不能小于等于 end，只能小于 end, 否则可能 i + 1 时会数组越界
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
