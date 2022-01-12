package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type Stack struct {
	Elements []*Node
}

func flatten(root *Node) *Node {
	var output, le *Node

	s := Stack{}
	cp := root

	for cp != nil {
		currentNode := &Node{}
		currentNode.Val = cp.Val
		if cp.Child != nil {
			if cp.Next != nil {
				s.Push(cp.Next)
			}
			cp = cp.Child
		} else {
			cp = cp.Next
			if cp == nil && !s.IsEmpty() {
				cp = s.Pop()
			}
		}

		if output == nil {
			output = currentNode
			le = currentNode
		} else {
			currentNode.Prev = le
			le.Next = currentNode
			le = le.Next
		}
	}
	return output
}

func (s *Stack) Push(input *Node) {
	s.Elements = append(s.Elements, input)
}

func (s *Stack) Pop() (output *Node) {
	output = s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) <= 0
}
