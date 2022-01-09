package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	h := head
	if h == nil {
		return nil
	}
	for h.Next != nil {
		t := h.Next
		h.Next = h.Next.Next
		t.Next = head
		head = t
	}
	return head
}
