/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (42.35%)
 * Likes:    2984
 * Dislikes: 108
 * Total Accepted:    404.1K
 * Total Submissions: 950.5K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 * 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 * 
 * Output: 3
 * 
 */
 type T struct {
	m int
	n int
	direction [][]int // 前后左右四个方向
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var obj T
	obj.m = len(grid)
	obj.n = len(grid[0])
	// 前后左右
	obj.direction = [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{-1, 0},
		[]int{1, 0},
	}
	var num = 0
	for i := 0; i < obj.m; i++ {
		for j := 0; j < obj.n; j++ {
			if grid[i][j] != '0' {
				obj.dfs(grid, i, j)
				num++	
			}
		}
	}
	return num
}

func (t T) dfs (grid [][]byte, r, c int) {
	if r < 0 || r >= t.m || c <0 || c >= t.n || grid[r][c] == 0 {
		return
	}
	grid[r][c] = '0' // 访问过的节点置零 -- 也就是一片联通的岛最终只会算一个岛
	for _, d := range t.direction {
		t.dfs(grid, r + d[0], c + d[1]);
	}
	return
}

