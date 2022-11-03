package link

func removeElements(head *ListNode, val int) *ListNode {
	h := new(ListNode)
	h.Next = head
	pre := h
	cur := h.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return h.Next
}
