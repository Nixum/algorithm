package link

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := new(ListNode)
	h.Next = head
	p := head
	for n > 0 && p != nil {
		p = p.Next
		n--
	}
	if p == nil && n == 0 {
		return head.Next
	}
	if p == nil {
		return nil
	}
	q := h
	for p != nil {
		q = q.Next
		p = p.Next
	}
	q.Next = q.Next.Next
	return h.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	h := new(ListNode)
	h.Next = head
	slow := h
	fast := head
	for n > 0 && fast != nil {
		n--
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return h.Next
}