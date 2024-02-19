package lc008_mid

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Tips: 这道题官方界定难度为mid，但感觉应该为easy级别，思路和实现都很简单
时间复杂度：O(max(m,n))，其中 m 和 n 分别为两个链表的长度。我们要遍历两个链表的全部位置，而处理每个位置只需要 O(1) 的时间。
空间复杂度：O(1)注意返回值不计入空间复杂度。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry

		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}

	}

	//如果链表遍历结束后，有 carry>0，还需要在答案链表的后面附加一个节点，节点的值为 carry. 即进位后的最高位
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
