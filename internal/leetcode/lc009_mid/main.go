package lc009_mid

/*
*
深度优先搜索
时间复杂度：O(MN)，其中 M 和 N 分别为行数和列数。
空间复杂度：O(MN)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 MN。
*/
func numIslands(grid [][]byte) int {
	res := 0

	if len(grid) == 0 {
		return 0
	}

	nr, nc := len(grid), len(grid[0])

	var dfs func(grid [][]byte, r, c int)
	dfs = func(grid [][]byte, r, c int) {
		nr = len(grid)
		nc = len(grid[0])

		if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				res++
				dfs(grid, r, c)
			}
		}
	}

	return res
}

/**
广度优先搜索
时间复杂度：O(MN，其中 M 和 N 分别为行数和列数。
空间复杂度：O(min(M,N))，在最坏情况下，整个网格均为陆地，队列的大小可以达到 min(M,N)。
*/
