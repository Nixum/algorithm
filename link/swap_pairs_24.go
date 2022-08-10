package link

func swapPairs(h *ListNode) *ListNode {
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	head := new(ListNode)
	pre := head
	p := h
	for p != nil {
		if p.Next != nil {
			pre.Next = p.Next
			tmp := p.Next.Next
			p.Next.Next = p
			p.Next = tmp
		}
		pre = p
		p = p.Next
	}
	return head.Next
}

func swapPairs2(h *ListNode) *ListNode {
	head := new(ListNode)
	head.Next = h
	pre := head
	for h != nil && h.Next != nil {
		pre.Next = h.Next
		tmp := h.Next.Next
		h.Next.Next = h
		h.Next = tmp
		pre = h
		h = h.Next
	}
	return head.Next
}