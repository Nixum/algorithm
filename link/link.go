package link

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func NewLink(arr []int) *ListNode {
	p := new(ListNode)
	h := p
	for _, i := range arr {
		p.Next = &ListNode{
			Val:  i,
		}
		p = p.Next
	}
	return h.Next
}

func Print(h *ListNode) {
	p := h
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}