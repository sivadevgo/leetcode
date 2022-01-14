package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	h := head
	tail := head
	var l int
	for h != nil {
		l++
		h = h.Next
		if h != nil {
			tail = h
		}
	}

	k = k % l
	k = l - k
	if k != 0 {
		tail.Next = head
		h = head

		for k > 1 {
			h = h.Next
			k--
		}
		head = h.Next
		h.Next = nil
	}

	return head

}
