package link

// 对链表从小到大排序
// 要求 O(nlogn)时间复杂度 和 O(1)空间复杂度
// 时间复杂度要求 O(nlogn)，可以用归并算法
// 由于要求空间复杂度是O(1)，所以用不了递归
// 所以使用自底向上的归并算法:
// 先两个两个的 merge，完成一趟后，再 4 个4个的 merge，直到结束。
// 例子：[4,3,1,7,8,9,2,11,5,6]
// step=1: (3->4)->(1->7)->(8->9)->(2->11)->(5->6)
// step=2: (1->3->4->7)->(2->8->9->11)->(5->6)
// step=4: (1->2->3->4->7->8->9->11)->5->6
// step=8: (1->2->3->4->5->6->7->8->9->11)
// 涉及到两个对链表的操作：
// merge(link1, link2): 双路归并
// cut(l, n)：将链表 l 切掉前 n 个节点，并返回后半部分的链表头
func sortList(head *ListNode) *ListNode {
	h := new(ListNode)
	h.Next = head
	p := head
	linkLen := 0
	for p != nil {
		p = p.Next
		linkLen++
	}
	for i := 1; i < linkLen; i *= 2 {
		cur := h.Next
		tail := h
		for cur != nil {
			left := cur
			right := cutLink(left, i)
			cur = cutLink(right, i)

			tail.Next = mergeLink(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return h.Next
}

// 切掉链表的前n个节点
func cutLink(head *ListNode, n int) *ListNode {
	h := head
	// 因为是从第一个开始算的，得先减
	n--
	for n > 0 && h != nil {
		h = h.Next
		n--
	}
	if h == nil {
		return nil
	}
	next := h.Next
	h.Next = nil
	return next
}

func mergeLink(l1 *ListNode, l2 *ListNode) *ListNode {
	h := new(ListNode)
	p := h
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return h.Next
}

// 归并排序，递归版本，时间复杂度要求 O(nlogn), 空间复杂度是 O(logn)
func sortList2(head *ListNode) *ListNode {
	// 这个判断很关键，否则会死循环
	if head == nil || head.Next == nil {
		return head
	}
	mid := returnLinkMidNode(head)
	left := sortList2(head)
	right := sortList2(mid)
	return mergeLink(left, right)
}

func returnLinkMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil
	return next
}
