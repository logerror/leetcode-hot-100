package lc022_easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
中序遍历，总是选择中间位置左边的数字作为根节点
时间复杂度：O(n)，其中 nnn 是数组的长度。每个数字只访问一次。
空间复杂度：O(logn)，其中 nnn 是数组的长度。空间复杂度不考虑返回值，因此空间复杂度主要取决于递归栈的深度，递归栈的深度是 O(logn)。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	var help func(nums []int, left, right int) *TreeNode
	help = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = help(nums, left, mid-1)
		root.Right = help(nums, mid+1, right)
		return root
	}

	return help(nums, 0, len(nums)-1)

}
