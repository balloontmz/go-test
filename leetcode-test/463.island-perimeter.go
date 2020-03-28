/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 *
 * https://leetcode.com/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (62.05%)
 * Likes:    1413
 * Dislikes: 97
 * Total Accepted:    164K
 * Total Submissions: 261.7K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * You are given a map in form of a two-dimensional integer grid where 1
 * represents land and 0 represents water.
 * 
 * Grid cells are connected horizontally/vertically (not diagonally). The grid
 * is completely surrounded by water, and there is exactly one island (i.e.,
 * one or more connected land cells).
 * 
 * The island doesn't have "lakes" (water inside that isn't connected to the
 * water around the island). One cell is a square with side length 1. The grid
 * is rectangular, width and height don't exceed 100. Determine the perimeter
 * of the island.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input:
 * [[0,1,0,0],
 * ⁠[1,1,1,0],
 * ⁠[0,1,0,0],
 * ⁠[1,1,0,0]]
 * 
 * Output: 16
 * 
 * Explanation: The perimeter is the 16 yellow stripes in the image below:
 * 
 * 
 * 
 * 
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var res int
    for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (grid[i][j] == 0) {continue}
            if (j == 0 || grid[i][j - 1] == 0) {res++}
            if (i == 0 || grid[i - 1][j] == 0) {res++}
            if (j == n - 1 || grid[i][j + 1] == 0) {res++}
            if (i == m - 1 || grid[i + 1][j] == 0) {res++}
		}
	}
	return res
}
// @lc code=end

