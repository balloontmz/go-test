/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 *
 * https://leetcode.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (58.15%)
 * Likes:    1202
 * Dislikes: 63
 * Total Accepted:    93.6K
 * Total Submissions: 160.3K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's
 * (representing land) connected 4-directionally (horizontal or vertical.) You
 * may assume all four edges of the grid are surrounded by water.
 * 
 * Find the maximum area of an island in the given 2D array. (If there is no
 * island, the maximum area is 0.)
 * 
 * Example 1:
 * 
 * 
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * 
 * Given the above grid, return 6. Note the answer is not 11, because the
 * island must be connected 4-directionally.
 * 
 * Example 2:
 * 
 * 
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 * 
 * Note: The length of each dimension in the given grid does not exceed 50.
 * 
 */
type T struct {
	m int
	n int
	direction [][]int // 前后左右四个方向
}

func maxAreaOfIsland(grid [][]int) int {
	// ✔ 726/726 cases passed (12 ms)
	// ✔ Your runtime beats 82.64 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (5.4 MB)
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
	var maxArea = 0
	for i := 0; i < obj.m; i++ {
		for j := 0; j < obj.n; j++ {
			maxArea = max(maxArea, obj.dfs(grid, i, j))
		}
	}
	return maxArea
}

func (t T) dfs (grid [][]int, r, c int) int {
	if r < 0 || r >= t.m || c <0 || c >= t.n || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0 // 访问过的节点置零
	var area = 1 // 清空该点后面积加 1
	for _, d := range t.direction {
		area += t.dfs(grid, r + d[0], c + d[1]);
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
