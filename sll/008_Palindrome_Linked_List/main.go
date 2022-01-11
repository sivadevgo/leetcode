package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}
	l := 0
	h := head
	for h != nil {
		l++
		h = h.Next
	}

	h = head
	lh1 := l / 2
	for i := 0; i < lh1; i++ {
		h = h.Next
	}
	if l%2 == 1 {
		h = h.Next
	}
	h2 := h

	for h2 != nil && h2.Next != nil {
		t := h2.Next
		h2.Next = h2.Next.Next
		t.Next = h
		h = t
	}
	h2 = head

	for i := 0; i < lh1; i++ {
		if h2 != nil && h != nil {
			if h2.Val != h.Val {
				return false
			}
		} else {
			return false
		}
		h2 = h2.Next
		h = h.Next
	}
	return true
}
