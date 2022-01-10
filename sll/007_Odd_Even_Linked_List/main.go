package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	var el *ListNode
	var ell *ListNode

	h := head

	for h != nil && h.Next != nil {
		if el == nil {
			el = h.Next
			ell = h.Next
		} else {
			ell.Next = h.Next
			ell = ell.Next
		}
		h.Next = h.Next.Next
		ell.Next = nil
		if h.Next != nil {
			h = h.Next
		}

	}
	// if h !=nil{
	h.Next = el
	// }

	return head
}
