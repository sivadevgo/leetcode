package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var output *ListNode
	var le *ListNode
	var reminder int

	for l1 != nil || l2 != nil {
		var newNode = &ListNode{}
		if l1 == nil && l2 != nil {
			newNode.Val, reminder = getModAndDivValue(l2.Val + reminder)
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			newNode.Val, reminder = getModAndDivValue(l1.Val + reminder)
			l1 = l1.Next
		} else {
			newNode.Val, reminder = getModAndDivValue(l1.Val + l2.Val + reminder)
			l1 = l1.Next
			l2 = l2.Next
		}

		if output == nil {
			output = newNode
			le = newNode
		} else {
			le.Next = newNode
			le = le.Next
		}
	}

	if reminder != 0 {
		var newNode = &ListNode{}
		newNode.Val = reminder
		le.Next = newNode
	}
	return output
}

func getModAndDivValue(input int) (modVal, divVal int) {
	return input % 10, input / 10
}
