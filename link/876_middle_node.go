package link

func middleNode(head *ListNode) *ListNode {
	h := new(ListNode)
	h.Next = head
	fast := h
	slow := h
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil && fast.Next == nil {
		return slow.Next
	}
	return slow
}

func middleNode2(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
