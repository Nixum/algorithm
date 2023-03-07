package link

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	link := NewLink([]int{1, 2, 6, 3, 4, 5, 6})
	Print(link)
	Print(removeElements(link, 6))
}

func TestReverseList(t *testing.T) {
	link := NewLink([]int{1, 2, 6, 3, 4, 5, 6})
	Print(reverseList(link))

	link = NewLink([]int{1})
	Print(reverseList(link))

	link = NewLink([]int{1, 2})
	Print(reverseList(link))
	fmt.Println()
	// 2--------
	link = NewLink([]int{1, 2, 6, 3, 4, 5, 6})
	Print(reverseList2(link))

	link = NewLink([]int{1})
	Print(reverseList2(link))

	link = NewLink([]int{1, 2})
	Print(reverseList2(link))
	fmt.Println()
	// 3----------------
	link = NewLink([]int{1, 2, 6, 3, 4, 5, 6})
	Print(reverseList3(link))

	link = NewLink([]int{1})
	Print(reverseList3(link))

	link = NewLink([]int{1, 2})
	Print(reverseList3(link))
}

func TestSwapPairs(t *testing.T) {
	link := NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	Print(swapPairs2(link))

	link = NewLink([]int{1, 2, 3, 4, 5, 6})
	Print(swapPairs2(link))

	link = NewLink([]int{1})
	Print(swapPairs2(link))

	link = NewLink([]int{})
	Print(swapPairs2(link))
}

func TestRemoveNthFromEnd(t *testing.T) {
	link := NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	Print(removeNthFromEnd2(link, 1))

	link = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	Print(removeNthFromEnd2(link, 5))

	link = NewLink([]int{1, 2, 3, 4, 5, 6})
	Print(removeNthFromEnd2(link, 6))
}

func TestGetIntersectionNode(t *testing.T) {
	var link1 *ListNode
	var link2 *ListNode
	link1 = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link2 = NewLink([]int{1, 2, 3})
	link2.Next.Next.Next = link1.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link2 = NewLink([]int{1})
	link2.Next = link1.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1})
	link2 = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link1.Next = link2.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))

	link1 = NewLink([]int{1, 2})
	link2 = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link1.Next.Next = link2.Next.Next.Next
	fmt.Println(getIntersectionNode(link1, link2))
}

func TestDetectCycle(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link.Next.Next.Next.Next.Next.Next.Next = link.Next.Next.Next
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1, 2, 3, 4})
	link.Next.Next.Next.Next = link.Next
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1})
	fmt.Println(detectCycle(link))

	link = NewLink([]int{1, 2})
	link.Next.Next = link
	fmt.Println(detectCycle(link))
}

func TestMergeTwoLists(t *testing.T) {
	var link1 *ListNode
	var link2 *ListNode
	link1 = NewLink([]int{1, 3, 5, 7})
	link2 = NewLink([]int{2, 4, 8})
	Print(mergeTwoLists(link1, link2))

	link1 = NewLink([]int{1, 2, 4})
	link2 = NewLink([]int{1, 3, 4})
	Print(mergeTwoLists(link1, link2))

	link1 = NewLink([]int{})
	link2 = NewLink([]int{})
	Print(mergeTwoLists(link1, link2))

	link1 = NewLink([]int{})
	link2 = NewLink([]int{0})
	Print(mergeTwoLists(link1, link2))
}

func TestHasCycle(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	link.Next.Next.Next.Next.Next.Next.Next = link.Next.Next.Next
	fmt.Println(hasCycle(link))

	link = NewLink([]int{1, 2, 3, 4})
	link.Next.Next.Next.Next = link.Next
	fmt.Println(hasCycle(link))

	link = NewLink([]int{1})
	fmt.Println(hasCycle(link))

	link = NewLink([]int{1, 2})
	link.Next.Next = link
	fmt.Println(hasCycle(link))

	link = NewLink([]int{1, 2})
	fmt.Println(hasCycle(link))
}

func TestMiddleNode(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(middleNode(link))

	link = NewLink([]int{1, 2, 3, 4})
	fmt.Println(middleNode(link))

	link = NewLink([]int{1})
	fmt.Println(middleNode(link))
}

func TestDeleteDuplicates(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 3, 3, 4, 5, 6, 6, 7})
	Print(deleteDuplicates(link))

	link = NewLink([]int{1, 1, 2})
	Print(deleteDuplicates(link))

	link = NewLink([]int{1, 1, 2, 3, 3})
	Print(deleteDuplicates(link))

	link = NewLink([]int{1, 2, 3})
	Print(deleteDuplicates(link))

	link = NewLink([]int{1, 1, 1})
	Print(deleteDuplicates(link))
}

func TestIsPalindrome(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(isPalindrome(link))

	link = NewLink([]int{1, 2, 3, 3, 2, 1})
	fmt.Println(isPalindrome(link))

	link = NewLink([]int{1, 2, 3, 5, 3, 2, 1})
	fmt.Println(isPalindrome(link))

	link = NewLink([]int{1, 2})
	fmt.Println(isPalindrome(link))
}

func TestReverseBetween(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseBetween(link, 2, 4))

	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseBetween(link, 1, 3))

	link = NewLink([]int{1})
	Print(reverseBetween(link, 1, 1))

	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseBetween(link, 4, 5))

	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseBetween(link, 1, 5))
}

func TestReverseKGroup(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseKGroup(link, 2))

	link = NewLink([]int{1, 2, 3, 4, 5})
	Print(reverseKGroup(link, 3))
}

func TestMergeKLists(t *testing.T) {
	var link1 *ListNode
	var link2 *ListNode
	var link3 *ListNode
	link1 = NewLink([]int{1, 4, 5})
	link2 = NewLink([]int{1, 3, 4})
	link3 = NewLink([]int{2, 6})
	Print(mergeKLists([]*ListNode{link1, link2, link3}))
}

func TestSortList(t *testing.T) {
	var link *ListNode
	link = NewLink([]int{3, 2, 5, 1, 4})
	Print(sortList2(link))

	//link = NewLink([]int{3, 2, 5, 1})
	//Print(sortList2(link))
}
