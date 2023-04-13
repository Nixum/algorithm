package link

// 将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，
// 然后返回重新排序的列表，需要保持相对顺序，
// 第一个节点索引为奇数，第二个为偶数
// 要求O(n)时间复杂度和O(1)空间复杂度
func oddEvenList(head *ListNode) *ListNode {
	n := head
	odd := &ListNode{}
	even := &ListNode{}
	var oddNext *ListNode
	var evenNext *ListNode
	// 每次处理两个
	for n != nil {
		tmp := n.Next
		// 处理奇数
		if odd.Next == nil {
			odd.Next = n
			oddNext = n
		} else {
			oddNext.Next = n
			oddNext = oddNext.Next
		}
		// 处理偶数
		if even.Next == nil {
			even.Next = n.Next
			evenNext = n.Next
		} else {
			evenNext.Next = n.Next
			evenNext = evenNext.Next
		}

		if tmp != nil {
			n = tmp.Next
		} else {
			n = tmp
		}
	}
	oddNext.Next = even.Next
	return odd.Next
}

// 直观点的写法
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h2 := head.Next
	odd := head
	even := h2
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = h2
	return head
}