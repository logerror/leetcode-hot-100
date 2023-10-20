package lc011_easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	//反转
	var head2 *ListNode = nil
	for slow != nil {
		tmp := slow.Next
		slow.Next = head2
		head2 = slow
		slow = tmp
	}

	//比对
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTemp
	}
	return prev
}
