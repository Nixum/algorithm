package link

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	link := NewLink([]int{1,2,6,3,4,5,6})
	Print(link)
	Print(removeElements(link, 6))
}

func TestReverseList(t *testing.T) {
	link := NewLink([]int{1,2,6,3,4,5,6})
	Print(reverseList(link))

	link = NewLink([]int{1})
	Print(reverseList(link))

	link = NewLink([]int{1, 2})
	Print(reverseList(link))
	// 2--------
	link = NewLink([]int{1,2,6,3,4,5,6})
	Print(reverseList2(link))

	link = NewLink([]int{1})
	Print(reverseList2(link))

	link = NewLink([]int{1, 2})
	Print(reverseList2(link))
}

func TestSwapPairs(t *testing.T) {
	link := NewLink([]int{1,2,3,4,5,6,7})
	Print(swapPairs2(link))

	link = NewLink([]int{1,2,3,4,5,6})
	Print(swapPairs2(link))

	link = NewLink([]int{1})
	Print(swapPairs2(link))

	link = NewLink([]int{})
	Print(swapPairs2(link))
}

func TestRemoveNthFromEnd(t *testing.T) {
	link := NewLink([]int{1,2,3,4,5,6,7})
	Print(removeNthFromEnd2(link, 1))

	link = NewLink([]int{1,2,3,4,5,6,7})
	Print(removeNthFromEnd2(link, 5))

	link = NewLink([]int{1,2,3,4,5,6})
	Print(removeNthFromEnd2(link, 6))
}

func TestGetIntersectionNode(t *testing.T) {
	var link1 *ListNode
	var link2 *ListNode
	link1 = NewLink([]int{1,2,3,4,5,6,7})
	link2 = NewLink([]int{1,2,3})
	link2.Next.Next.Next = link1.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1,2,3,4,5,6,7})
	link2 = NewLink([]int{1})
	link2.Next = link1.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1})
	link2 = NewLink([]int{1,2,3,4,5,6,7})
	link1.Next = link2.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1, 2})
	link2 = NewLink([]int{1,2,3,4,5,6,7})
	link1.Next.Next = link2.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))
}

func TestDetectCycle(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1,2,3,4,5,6,7})
	link.Next.Next.Next.Next.Next.Next.Next = link.Next.Next.Next
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1,2,3,4})
	link.Next.Next.Next.Next = link.Next
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1})
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1, 2})
	link.Next.Next = link
	fmt.Println(detectCycle(link))
}