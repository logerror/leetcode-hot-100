package lc013_mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*

	中序遍历

二叉搜索树具有如下性质：

	结点的左子树只包含小于当前结点的数。
	结点的右子树只包含大于当前结点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

二叉树的中序遍历即按照访问左子树——根结点——右子树的方式遍历二叉树；在访问其左子树和右子树时，我们也按照同样的方式遍历；直到遍历完整棵树。

思路和算法

	因为二叉搜索树和中序遍历的性质，所以二叉搜索树的中序遍历是按照键增加的顺序进行的。于是，我们可以通过中序遍历找到第 kkk 个最小元素。
	「二叉树的中序遍历」可以参考「94. 二叉树的中序遍历的官方题解」，具体地，我们使用迭代方法，这样可以在找到答案后停止，不需要遍历整棵树。

时间复杂度：O(H+k)，其中 HHH 是树的高度。在开始遍历之前，我们需要 O(H)到达叶结点。当树是平衡树时，时间复杂度取得最小值 O(log⁡N+k)；当树是线性树（树中每个结点都只有一个子结点或没有子结点）时，时间复杂度取得最大值 O(N+k)。

空间复杂度：O(H)，栈中最多需要存储 H 个元素。当树是平衡树时，空间复杂度取得最小值 O(log⁡N)；当树是线性树时，空间复杂度取得最大值 O(N)
*/
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// 记录子树的结点数
/**
在方法一中，我们之所以需要中序遍历前 kkk 个元素，是因为我们不知道子树的结点数量，不得不通过遍历子树的方式来获知。
因此，我们可以记录下以每个结点为根结点的子树的结点数，并在查找第 kkk 小的值时，使用如下方法搜索：
令 node 等于根结点，开始搜索。
对当前结点 node进行如下操作：

如果 node 的左子树的结点数 left 小于 k−1，则第 k 小的元素一定在 node 的右子树中，令 node 等于其的右子结点，k 等于 k−left−1，并继续搜索；
如果 node 的左子树的结点数 left 等于 k−1，则第 k 小的元素即为 node ，结束搜索并返回 node 即可；
如果 node 的左子树的结点数 left 大于 k−1，则第 k 小的元素一定在 node 的左子树中，令 node 等于其左子结点，并继续搜索。
在实现中，我们既可以将以每个结点为根结点的子树的结点数存储在结点中，也可以将其记录在哈希表中。
*/

type MyBst struct {
	root *TreeNode
	// 统计以每个节点为根节点的子树的节点数
	nodeNum map[*TreeNode]int
}

// 统计自身及子树节点数
func (t *MyBst) countNodeNums(node *TreeNode) int {
	if node == nil {
		return 0
	}
	t.nodeNum[node] = 1 + t.countNodeNums(node.Left) + t.countNodeNums(node.Right)
	return t.nodeNum[node]
}

func (t *MyBst) kthSmallest2(k int) int {
	node := t.root
	for {
		leftNodeNum := t.nodeNum[node.Left]
		if leftNodeNum < k-1 {
			node = node.Right
			k = k - leftNodeNum + 1
		} else if leftNodeNum == k-1 {
			return node.Val
		} else {
			node = node.Left
		}
	}
}

func kthSmallest2(root *TreeNode, k int) int {
	t := MyBst{root, map[*TreeNode]int{}}
	t.countNodeNums(root)
	return t.kthSmallest2(k)
}
