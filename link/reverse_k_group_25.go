package link

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseKGroupBetween(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 左闭右开区间
func reverseKGroupBetween(start *ListNode, end *ListNode) *ListNode {
	p := start
	var pre *ListNode
	for p != end {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

// 错误的思路，这样会导致最后不足k的链表没有接上
func bad_reverseKGroup2(head *ListNode, k int) *ListNode {
	p := head
	hLen := 0
	for p != nil {
		p = p.Next
		hLen ++
	}
	num := hLen / k
	p = head
	start, nextStart := reverseKGroupByNum2(p, k)
	h := start
	for i := 1; i <= num; i++ {
		h1 := start
		for h1 != nil {
			if h1.Next == nil {
				start, nextStart = reverseKGroupByNum2(nextStart, k)
				break
			}
			h1 = h1.Next
		}
		if i == num {
			for h1 != nil {
				if h1.Next == nil {
					h1.Next = start
					break
				}
			}
			break
		}
	}
	return h
}

func reverseKGroupByNum2(head *ListNode, k int) (*ListNode, *ListNode) {
	hLen := 0
	h := head
	for h != nil {
		hLen++
		h = h.Next
	}
	if hLen < k {
		return head, nil
	}
	var pre *ListNode
	p := head
	for k > 0 {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
		k--
	}
	return pre, p
}

// 解决不了当head不足k个的时候，不进行反转
func reverseKGroupByNum(head *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode
	p := head
	for k > 0 {
		if p == nil {
			// 此时head链已经被修改了
			return head, nil
		}
		next := p.Next
		p.Next = pre
		pre = p
		p = next
		k--
	}
	return pre, p
}

// 错误的思路, 边遍历边换，导致整条链乱了
func bad_reverseKGroup3(head *ListNode, k int) *ListNode {
	p := head
	hLen := 1
	for p != nil {
		p = p.Next
		hLen ++
	}
	num := hLen / k
	var pre *ListNode
	p = head
	curK := 0
	for num != 0 {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
		curK++
		if curK == k {
			tmpPre := pre
			for tmpPre != nil {
				tmpPre = tmpPre.Next
				if tmpPre.Next == nil {
					tmpPre.Next = p
					pre = nil
					break
				}
			}
			num--
			curK = 0
		}
		if num == 0 {
			for pre != nil {
				pre = pre.Next
				if pre.Next == nil {
					pre.Next = p
					break
				}
			}
			break
		}
	}
	return head
}
