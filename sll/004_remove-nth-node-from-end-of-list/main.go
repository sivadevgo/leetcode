package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil && head.Next == nil {
		return nil
	}
	h1 := head
	var c int

	for h1 != nil {
		c++
		h1 = h1.Next
	}

	r := c - n - 1
	if r == -1 {
		head = head.Next
		return head
	}
	h1 = head
	for r > 0 {
		r--
		h1 = h1.Next
	}
	if h1.Next != nil {
		h1.Next = h1.Next.Next
	}
	return head
}
