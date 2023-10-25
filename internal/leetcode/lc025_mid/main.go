package lc025_mid

func searchMatrix(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])

	for i := 0; i < row; i++ {
		// 当前数字
		if matrix[i][column-1] == target {
			return true
		}

		// 在下一行
		if matrix[i][column-1] < target {
			continue
		}

		// 在当前行
		if matrix[i][column-1] > target {
			for j := 0; j < column-1; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}

	}
	return false
}
