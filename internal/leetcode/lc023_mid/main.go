package lc023_mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的层序遍历，取每层的最右侧节点值
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, q[len(q)-1].Val)
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
