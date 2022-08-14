package link

func deleteDuplicates(head *ListNode) *ListNode {
	h := new(ListNode)
	h.Val = -101
	h.Next = head
	pre := h
	p := head
	for p != nil {
		if p.Val == pre.Val {
			pre.Next = p.Next
		} else {
			pre = p
		}
		p = p.Next
	}
	return h.Next
}
