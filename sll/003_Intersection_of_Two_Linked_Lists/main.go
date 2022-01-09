package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	h1, h2 := headA, headB
	var cnt int

	for h1 != nil {
		cnt++
		h1 = h1.Next
	}

	for h2 != nil {
		cnt--
		h2 = h2.Next
	}

	h1, h2 = headA, headB

	if cnt > 0 {
		for cnt > 0 {
			cnt--
			h1 = h1.Next
		}
	} else if cnt < 0 {
		for cnt < 0 {
			cnt++
			h2 = h2.Next
		}
	}

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h1
		} else {
			h1 = h1.Next
			h2 = h2.Next
		}
	}
	return nil
}
