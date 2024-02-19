package lc007_mid

/**
方法一：模拟
可以模拟螺旋矩阵的路径。初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。

判断路径是否进入之前访问过的位置需要使用一个与输入矩阵大小相同的辅助矩阵 visited\textit{visited}visited，其中的每个元素表示该位置是否被访问过。当一个元素被访问时，将 visited\textit{visited}visited 中的对应位置的元素设为已访问。

如何判断路径是否结束？由于矩阵中的每个元素都被访问一次，因此路径的长度即为矩阵中的元素数量，当路径的长度达到矩阵中的元素数量时即为完整路径，将该路径返回。
*/

func spiralOrder(matrix [][]int) []int {
	// 为空直接返回
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])

	//初始化数组用于记录某个元素是否被访问
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total       = rows * columns
		order       = make([]int, total)
		row, column = 0, 0
		// 对应四个方向>
		//行数不变，列+1
		//列数不变，行+1
		//行数不变，列-1
		//列数不变，行-1
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		//标记元素已经被访问
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}

	return order
}

/*
*
可以将矩阵看成若干层，首先输出最外层的元素，其次输出次外层的元素，直到输出最内层的元素。

定义矩阵的第 kkk 层是到最近边界距离为 kkk 的所有顶点。例如，下图矩阵最外层元素都是第 1 层，次外层元素都是第 22 层，剩下的元素都是第 3 层。

[[1, 1, 1, 1, 1, 1, 1],

	[1, 2, 2, 2, 2, 2, 1],
	[1, 2, 3, 3, 3, 2, 1],
	[1, 2, 2, 2, 2, 2, 1],
	[1, 1, 1, 1, 1, 1, 1]]

对于每层，从左上方开始以顺时针的顺序遍历所有元素。假设当前层的左上角位于 (top,left)，右下角位于 (bottom,right)，按照如下顺序遍历当前层的元素。

从左到右遍历上侧元素，依次为 (top,left) 到 (top,right)。

从上到下遍历右侧元素，依次为 (top+1,right) 到 (bottom,right)

如果 left<right且 top<bottom，则从右到左遍历下侧元素，依次为 (bottom,right−1) 到 (bottom,left+1)，以及从下到上遍历左侧元素，依次为 (bottom,left)到 (top+1,left)

遍历完当前层的元素之后，将 left 和 top 分别增加 111，将 right 和 bottom分别减少 111，进入下一层继续遍历，直到遍历完所有元素为止。

作者：力扣官方题解
链接：https://leetcode.cn/problems/spiral-matrix/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
