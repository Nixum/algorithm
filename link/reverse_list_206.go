package link

func reverseList(head *ListNode) *ListNode {
	h := new(ListNode)
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = h.Next
		h.Next = p
		p = tmp
	}
	return h.Next
}

// 区别上面那种写法，无需头节点，直接在原链表上操作
// 但是会破坏head节点
func reverseList3(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
