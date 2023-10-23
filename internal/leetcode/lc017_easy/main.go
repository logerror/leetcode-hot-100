package lc017_easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// 利用哈希表
/*
复杂度分析

时间复杂度：O(N)，其中 N 是链表中的节点数。最坏情况下我们需要遍历每个节点一次。

空间复杂度：O(N)，其中 N 是链表中的节点数。主要为哈希表的开销，最坏情况下我们需要将每个节点插入到哈希表中一次。
*/
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

//快慢指针
/*
时间复杂度：O(N)，其中 N 是链表中的节点数。

当链表中不存在环时，快指针将先于慢指针到达链表尾部，链表中每个节点至多被访问两次。

当链表中存在环时，每一轮移动后，快慢指针的距离将减小一。而初始距离为环的长度，因此至多移动 N 轮。

空间复杂度：O(1)。我们只使用了两个指针的额外空间。

*/

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
