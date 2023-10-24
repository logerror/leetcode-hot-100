package lc021_easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归
假设树上一共 n 个节点。

时间复杂度：这里遍历了这棵树，渐进时间复杂度为 O(n)。
空间复杂度：这里的空间复杂度和递归使用的栈空间有关，这里递归层数不超过 nnn，故渐进空间复杂度为 O(n)
*/
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

/*
迭代
复杂度分析

时间复杂度：O(n)，同「方法一」。
空间复杂度：这里需要用一个队列来维护节点，每个节点最多进队一次，出队一次，队列中最多不会超过 n 个点，故渐进空间复杂度为 O(n)。

*/

func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)

	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}

		if u.Val != v.Val {
			return false
		}

		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left)
	}

	return true
}
