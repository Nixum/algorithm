package link

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aLen := 0
	bLen := 0
	p := headA
	q := headB
	for p != nil {
		p = p.Next
		aLen++
	}
	for q != nil {
		q = q.Next
		bLen++
	}
	n := 0
	if aLen > bLen {
		p, q, n = headA, headB, aLen - bLen
	} else {
		p, q, n = headB, headA, bLen - aLen
	}
	for n > 0 {
		p = p.Next
		n--
	}
	for p != nil {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}
	return nil
}
