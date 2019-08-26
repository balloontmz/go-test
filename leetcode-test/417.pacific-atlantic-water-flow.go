/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 *
 * https://leetcode.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (38.00%)
 * Likes:    743
 * Dislikes: 134
 * Total Accepted:    49.3K
 * Total Submissions: 129.2K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * Given an m x n matrix of non-negative integers representing the height of
 * each unit cell in a continent, the "Pacific ocean" touches the left and top
 * edges of the matrix and the "Atlantic ocean" touches the right and bottom
 * edges.
 * 
 * Water can only flow in four directions (up, down, left, or right) from a
 * cell to another one with height equal or lower.
 * 
 * Find the list of grid coordinates where water can flow to both the Pacific
 * and Atlantic ocean.
 * 
 * Note:
 * 
 * 
 * The order of returned grid coordinates does not matter.
 * Both m and n are less than 150.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given the following 5x5 matrix:
 * 
 * ⁠ Pacific ~   ~   ~   ~   ~ 
 * ⁠      ~  1   2   2   3  (5) *
 * ⁠      ~  3   2   3  (4) (4) *
 * ⁠      ~  2   4  (5)  3   1  *
 * ⁠      ~ (6) (7)  1   4   5  *
 * ⁠      ~ (5)  1   1   2   4  *
 * ⁠         *   *   *   *   * Atlantic
 * 
 * Return:
 * 
 * [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with
 * parentheses in above matrix).
 * 
 * 
 * 
 * 
 */
var m, n int
var gMatrix [][]int
var direction = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

func pacificAtlantic(matrix [][]int) [][]int {
	// ✔ 113/113 cases passed (40 ms)
	// ✔ Your runtime beats 96.88 % of golang submissions
	// ✔ Your memory usage beats 66.67 % of golang submissions (22.3 MB)
	var ret = [][]int{}
	if matrix == nil || len(matrix) == 0 {
		return ret
	}
	m = len(matrix)
	n = len(matrix[0])
	gMatrix = matrix
	// 分别用于记录所有点是否能到达 太平洋和大西洋
	var canReachP = make([][]bool, m)
	var canReachA = make([][]bool, m)
	for i := 0; i < m; i++ {
		canReachP[i] = make([]bool, n)
		canReachA[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, canReachP)
		dfs(i, n-1, canReachA)
	}

	for i := 0; i < n; i++ {
		dfs(0, i, canReachP)
		dfs(m-1, i, canReachA)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[i][j] && canReachA[i][j] {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret
}

// 判断某个点是否可以到达某个洋，如果可以，则将对应的数组值标记为 true
func dfs(r, c int, canReach [][]bool)  {
	if canReach[r][c] { // 此处为 true 代表此处已经进入过，此点后面的节点也就没有了遍历的必要
		return
	}
	canReach[r][c] = true
	for _, d := range direction {
		var nextR = d[0] + r
		var nextC = d[1] + c
		// 此处由于要做两个节点的比较，所以并未放在最外面，而是放在方向遍历内
		if nextR < 0 || nextR >= m || nextC < 0 || nextC >= n || gMatrix[r][c] > gMatrix[nextR][nextC] {
			continue
		}
		dfs(nextR, nextC, canReach)
	}
}