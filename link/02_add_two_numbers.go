package link

// 正序写法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resHead := &ListNode{}
	carry := 0
	var pre *ListNode
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		t := carry + v1 + v2
		carry = t / 10
		n := &ListNode{
			Val:  t%10,
			Next: nil,
		}
		if resHead.Next == nil {
			resHead.Next = n
			pre = n
		} else {
			pre.Next = n
			pre = n
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil && carry > 0 {
			pre.Next = &ListNode{
				Val:  carry,
			}
		}
	}
	return resHead.Next
}

// 倒序写法，后面还要翻转链表
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	resHead := &ListNode{}
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		t := carry + v1 + v2
		carry = t / 10
		n := &ListNode{
			Val:  t%10,
			Next: nil,
		}
		n.Next = resHead.Next
		resHead.Next = n
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		tmp := resHead.Next
		resHead.Next = &ListNode{
			Val:  carry,
			Next: tmp,
		}
	}
	return reverseInAddTwoNumbers(resHead).Next
}

func reverseInAddTwoNumbers(head *ListNode) *ListNode {
	h := &ListNode{}
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = h.Next
		h.Next = p
		p = tmp
	}
	return h.Next
}
