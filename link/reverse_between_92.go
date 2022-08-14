package link

// 递归解法, 下面那种解法的简便写法
func simpleReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseBetweenEnd(head, right)
	}
	head.Next = reverseBetween(head.Next, left - 1, right - 1)
	return head
}

// 递归解法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	lastStart = nil
	p := head
	var pre *ListNode
	leftIndex := 1
	for leftIndex != left {
		leftIndex++
		pre = p
		p = p.Next
	}
	if pre != nil {
		pre.Next = reverseBetweenEnd(p, right - left + 1)
	} else {
		return reverseBetweenEnd(p, right - left + 1)
	}
	return head
}

var lastStart *ListNode
func reverseBetweenEnd(head *ListNode, num int) *ListNode {
	if num == 1 {
		lastStart = head.Next
		return head
	}
	last := reverseBetweenEnd(head.Next, num - 1)
	head.Next.Next = head
	head.Next = lastStart
	return last
}

// 非递归解法
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	p := head
	leftIndex := 1
	for leftIndex != left {
		leftIndex++
		p = p.Next
	}
	num := right - left + 1
	var pre *ListNode
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
		num--
		if num == 0 {
			// 拼上尾部
			h := pre
			for h != nil {
				if h.Next == nil {
					h.Next = p
					break
				}
				h = h.Next
			}
			break
		}
	}
	res := new(ListNode)
	res.Next = head
	p2 := res
	for p2 != nil {
		left--
		if left == 0 {
			p2.Next = pre
			break
		}
		p2 = p2.Next
	}
	return res.Next
}

// 错误的答案，因为修改了head的缘故，导致拼不起来
func bad_ReverseBetween(head *ListNode, left int, right int) *ListNode {
	h := new(ListNode)
	h.Next = head
	pre := h
	p := head
	leftIndex := 1
	if leftIndex != left {
		leftIndex++
		p = p.Next
		pre = pre.Next
	}
	num := right - left + 1
	h1 := new(ListNode)
	var last *ListNode
	for p != nil {
		next := p.Next
		// 这里会导致断开
		p.Next = h1.Next
		h1.Next = p
		p = next
		num--
		if num == 0 {
			h2 := h1
			for h2 != nil {
				if h2.Next == nil {
					h2.Next = last
					break
				}
				h2 = h2.Next
			}
			break
		}
	}
	pre.Next = h1.Next
	res := new(ListNode)
	res.Next = head
	p2 := res
	for p2 != nil {
		if p2.Next == pre {
			p2.Next = pre
			continue
		}
		p2 = p2.Next
	}
	return res.Next
}
