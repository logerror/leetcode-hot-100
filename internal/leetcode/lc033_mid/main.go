package lc033_mid

//单指针
/*
我们可以考虑对数组进行两次遍历。在第一次遍历中，我们将数组中所有的 0 交换到数组的头部。在第二次遍历中，我们将数组中所有的 1 交换到头部的 0 之后。此时，所有的 2 都出现在数组的尾部，这样我们就完成了排序。

具体地，我们使用一个指针 ptr 表示「头部」的范围，ptr 中存储了一个整数，表示数组 nums 从位置 0 到位置 ptr−1 都属于「头部」。ptr 的初始值为 0，表示还没有数处于「头部」。

在第一次遍历中，我们从左向右遍历整个数组，如果找到了 0，那么就需要将 0 与「头部」位置的元素进行交换，并将「头部」向后扩充一个位置。在遍历结束之后，所有的 0 都被交换到「头部」的范围，并且「头部」只包含 000。

在第二次遍历中，我们从「头部」开始，从左向右遍历整个数组，如果找到了 1，那么就需要将 1 与「头部」位置的元素进行交换，并将「头部」向后扩充一个位置。在遍历结束之后，所有的 1 都被交换到「头部」的范围，并且都在 0 之后，此时 2 只出现在「头部」之外的位置，因此排序完成。
*/
func sortColors(nums []int) {
	count0 := swapColors(nums, 0)
	swapColors(nums[count0:], 1)
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}
