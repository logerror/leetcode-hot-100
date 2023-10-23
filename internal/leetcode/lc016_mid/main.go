package lc016_mid

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
递归
时间复杂度：O(n)，其中n 是链表的节点数量。需要对每个节点进行更新指针的操作。

空间复杂度：O(n)，其中 n 是链表的节点数量。空间复杂度主要取决于递归调用的栈空间。
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := head.Next
	head.Next = swapPairs(head2.Next)
	head2.Next = head
	return head2
}

/*
迭代
时间复杂度：O(n)，其中 n 是链表的节点数量。需要对每个节点进行更新指针的操作。

空间复杂度：O(1)。
*/
func swapPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
