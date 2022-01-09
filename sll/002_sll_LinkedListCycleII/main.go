package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cycle bool
	m := map[*ListNode]bool{}
	first := head
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cycle = true
			break
		}
	}

	for cycle {
		if _, ok := m[slow]; !ok {
			m[slow] = true
			slow = slow.Next
		} else {
			break
		}
	}

	for first != nil {
		if _, ok := m[first]; ok {
			return first
		} else {
			first = first.Next
		}
	}
	return nil
}
