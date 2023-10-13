package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)

}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// [5,1,4,null,null,3,6]

// 中序遍历 时间复杂度 O(n) ，  空间复杂度 O(n)
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 取出树最左侧节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}

		inorder = root.Val
		root = root.Right
	}

	return true
}
