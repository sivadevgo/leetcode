package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var output, ln *ListNode
	for list1 != nil || list2 != nil {
		var nextNode *ListNode
		if list1 == nil && list2 != nil || list1 != nil && list2 != nil && list1.Val >= list2.Val {
			list2, nextNode = getFirstNode(list2)
		} else if list2 == nil && list1 != nil || list1 != nil && list2 != nil && list2.Val >= list1.Val {
			list1, nextNode = getFirstNode(list1)
		}
		if output == nil {
			output = nextNode
			ln = nextNode
		} else {
			ln.Next = nextNode
			ln = ln.Next
		}
	}
	return output
}

func getFirstNode(list *ListNode) (*ListNode, *ListNode) {
	nextNode := list
	list = list.Next
	nextNode.Next = nil
	return list, nextNode
}

//DOUBLY LINKED LIST
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     var output,ln *ListNode
//     for list1!=nil || list2!=nil{
//         var nextNode  *ListNode
//         if list1 == nil && list2!=nil || list1!=nil&&list2!=nil&&list1.Val >= list2.Val {
//             list2,nextNode = getFirstNode(list2)
//         }else if list2 == nil && list1!=nil || list1!=nil && list2 !=nil && list2.Val >= list1.Val{
//             list1,nextNode = getFirstNode(list1)
//         }
//         if output == nil{
//             output = nextNode
//             ln = nextNode
//         }else{
//             ln.Next = nextNode
//             nextNode.Prev = ln
//             ln = ln.Next
//         }
//     }
//     return output
// }

// func getFirstNode(list *ListNode)(list,element){
//     nextNode := list
//     list = list.Next
//     list.Prev = nil
//     nextNode.Next = nil
//     nextNode.Prev = nil
//     return list, nextNode
// }
