package link

// 另一种方案是使用优先级队列，遍历所有链表把所有节点加入优先级队列实现排序
// 两两合并链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var tmp *ListNode
	for i := 0; i < len(lists); i++ {
		tmp = mergeKListsByTwo(tmp, lists[i])
	}
	return tmp
}

func mergeKListsByTwo(h1 *ListNode, h2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for h1 != nil && h2 != nil {
		if h1.Val > h2.Val {
			p.Next = h2
			h2 = h2.Next
		} else {
			p.Next = h1
			h1 = h1.Next
		}
		p = p.Next
	}
	if h1 != nil {
		p.Next = h1
		h1 = h1.Next
		p = p.Next
	}
	if h2 != nil {
		p.Next = h2
		h2 = h2.Next
		p = p.Next
	}
	return head.Next
}
