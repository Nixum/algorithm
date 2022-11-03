package prefixsum

// 前缀和
type NumArray struct {
	res []int
}

func NumArrayConstructor(nums []int) NumArray {
	arr := NumArray{
		res: make([]int, len(nums) + 1),
	}
	for i := 1; i <= len(nums); i++ {
		arr.res[i] = arr.res[i - 1] + nums[i - 1]
	}
	return arr
}


func (n *NumArray) SumRange(left int, right int) int {
	return n.res[right + 1] - n.res[left]
}


// ----------------
type NumArray2 struct {
	res []int
}

func NumArrayConstructor2(nums []int) NumArray2 {
	arr := NumArray2{
		res: make([]int, len(nums)),
	}
	arr.res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr.res[i] = arr.res[i - 1] + nums[i]
	}
	return arr
}


func (n *NumArray2) SumRange(left int, right int) int {
	if left == 0 {
		return n.res[right]
	}
	return n.res[right] - n.res[left - 1]
}
