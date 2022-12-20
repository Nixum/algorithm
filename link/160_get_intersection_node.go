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

// 错误的思路，该方案解决不了，headA包含了headB的场景
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	stackA := make([]*ListNode, 0)
	stackB := make([]*ListNode, 0)
	ha := headA
	for ha != nil {
		stackA = append(stackA, ha)
		ha = ha.Next
	}
	hb := headB
	for hb != nil {
		stackB = append(stackB, hb)
		hb = hb.Next
	}
	if len(stackA) == 1 &&
		len(stackB) == len(stackB) &&
		stackB[0] == stackA[0]{
		return stackB[0]
	}
	aIndex := len(stackA) - 1
	bIndex := len(stackB) - 1
	for aIndex >= 0 && bIndex >= 0 {
		if stackA[aIndex] == stackB[bIndex] &&
			aIndex - 1 >= 0 &&
			bIndex - 1 >= 0 &&
			stackA[aIndex - 1] != stackB[bIndex - 1] {
			return stackA[aIndex]
		}
		aIndex--
		bIndex--
	}
	return nil
}