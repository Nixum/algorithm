package link

func deleteNode(node *ListNode) {
	for node.Next != nil && node.Next.Next != nil {
		node.Val, node.Next.Val = node.Next.Val, node.Val
		node = node.Next
	}
	node.Val, node.Next.Val = node.Next.Val, node.Val
	node.Next = nil
}
