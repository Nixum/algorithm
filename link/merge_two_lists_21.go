package link

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head1 := list1
	head2 := list2
	head := new(ListNode)
	p := head
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}
	for head1 != nil {
		p.Next = head1
		p = p.Next
		head1 = head1.Next
	}
	for head2 != nil {
		p.Next = head2
		p = p.Next
		head2 = head2.Next
	}
	return head.Next
}
