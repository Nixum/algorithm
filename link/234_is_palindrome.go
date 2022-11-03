package link

// 时间复杂度O(n), 空间复杂度O(1)
// 先用快慢指针找到中点，然后从中点开始反转链表 和 从头开始的节点开始比较
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := reverseLinkWhenIsPalindrome(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}

func reverseLinkWhenIsPalindrome(head *ListNode) *ListNode {
	//var pre *ListNode
	//cur := head
	//for cur != nil {
	//	tmp := cur.Next
	//	cur.Next = pre
	//	pre = cur
	//	cur = tmp
	//}
	//return pre
	h := new(ListNode)
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = h.Next
		h.Next = cur
		cur = tmp
	}
	return h.Next
}

// 但是这种空间复杂度是O(n)，不够优雅
func isPalindrome2(head *ListNode) bool {
	isPalindromeHead = head
	_, res := getNodeByReverse(head)
	isPalindromeHead = nil
	return res
}

var isPalindromeHead *ListNode
func getNodeByReverse(head *ListNode) (*ListNode, bool) {
	if head.Next == nil {
		return head, true
	}
	n, res := getNodeByReverse(head.Next)
	curRes := false
	if isPalindromeHead.Val == n.Val {
		curRes = true
	}
	isPalindromeHead = isPalindromeHead.Next
	return head, res && curRes
}
