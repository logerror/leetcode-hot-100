package lc011_easy

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {


}

func reverseList(head *ListNode) *ListNode{
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTemp
	}
	return prev
}