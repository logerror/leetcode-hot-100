package lc001

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lLen := dfs(node.Left)
		rLen := dfs(node.Right)
		ans = max(ans, lLen+rLen+1)
		return max(lLen, rLen) + 1
	}
	dfs(root)
	return ans - 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
