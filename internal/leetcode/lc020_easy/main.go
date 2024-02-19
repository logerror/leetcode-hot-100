package lc020_easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, node1, node2 := dummyHead, list1, list2
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			temp.Next = node1
			node1 = node1.Next
		} else {
			temp.Next = node2
			node2 = node2.Next
		}
		temp = temp.Next
	}

	if node1 != nil {
		temp.Next = node1
	}

	if node2 != nil {
		temp.Next = node2
	}

	return dummyHead.Next
}
