/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (23.52%)
 * Likes:    883
 * Dislikes: 460
 * Total Accepted:    157.8K
 * Total Submissions: 665.6K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 * 
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 * 
 * Example:
 * 
 * 
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 * 
 * 
 * After running your function, the board should be:
 * 
 * 
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 * 
 * 
 * Explanation:
 * 
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 * 
 */
var direction = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

var m, n int

func solve(board [][]byte)  {
	// ✔ 59/59 cases passed (36 ms)
	// ✔ Your runtime beats 75.34 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (9.2 MB)
	if board == nil || len(board) == 0 {
		return
	}
	m = len(board)
	n = len(board[0])

	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// 对没被包围的节点进行深度优先遍历，同时替换所有未被包围的未中间 byte T
// 被包围的节点 无法进入此遍历
func dfs(board [][]byte, r, c int)  {
	if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != 'O' {
		return 
	}
	board[r][c] = 'T'
	for _, d := range direction {
		dfs(board, r + d[0], c + d[1])
	}
}
