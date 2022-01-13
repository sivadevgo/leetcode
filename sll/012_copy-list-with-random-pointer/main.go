package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	h := head
	var newHead *Node
	var newTail *Node
	// vm := map[int]*Node{}
	im := map[int]*Node{}
	pm := map[*Node]int{}
	var i int

	for h != nil {
		newNode := &Node{}
		newNode.Val = h.Val
		// vm[newNode.Val] = newNode
		im[i] = newNode
		pm[h] = i
		if newHead == nil {
			newHead = newNode
			newTail = newHead
		} else {
			newTail.Next = newNode
			newTail = newTail.Next
		}
		h = h.Next
		i++
	}

	h = head
	nh := newHead
	for h != nil {
		r := h.Random
		if r != nil {
			index, ok := pm[r]
			if ok {
				node, _ := im[index]
				nh.Random = node
			}
			// rn,ok:=vm[r.Val]
			// if ok{
			//     nh.Random = rn
			// }
		}

		h = h.Next
		nh = nh.Next
	}
	return newHead
}
