/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (31.95%)
 * Likes:    2158
 * Dislikes: 109
 * Total Accepted:    323.8K
 * Total Submissions: 1M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 * 
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 * 
 * Example:
 * 
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 * 
 * 
 */

type Obj struct {
	direction [][]int
	m int
	n int
	board [][]byte
	word string
}

func exist(board [][]byte, word string) bool {
	// ✔ 87/87 cases passed (8 ms)
	// ✔ Your runtime beats 47.88 % of golang submissions
	// ✔ Your memory usage beats 50 % of golang submissions (3.6 MB)
    if len(word) == 0 {
		return true
	}
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	var obj Obj

	obj.direction = [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
	}
	obj.m = len(board)
	obj.n = len(board[0])
	var hasVisited = make([][]bool, obj.m)
	for i := 0; i < obj.m; i++ {
		hasVisited[i] = make([]bool, obj.n)
	}

	for r := 0; r < obj.m; r++ {
		for c := 0; c < obj.n; c++ {
			if obj.backtracking(0, r, c, hasVisited, board, word) {
				return true
			}
		}
	}
	return false
}

func (obj Obj) backtracking(curLen, r, c int, visited [][]bool, board [][]byte, word string) bool {
	if curLen == len(word) {
		return true
	}
	// 如果超过数组限度，或者当前递归的节点不等于当前遍历次数的rune，或者当前节点已经遍历
	// rune -- byte 和 rune 的比较？
	if r < 0 || r >= obj.m || c < 0 || c >= obj.n || board[r][c] != word[curLen] || visited[r][c] {
		return false	
	}
	visited[r][c] = true
	for _, v := range obj.direction {
		if obj.backtracking(curLen + 1, r + v[0], c + v[1], visited, board, word) {
			return true	
		}
	}
	visited[r][c] = false
	return false
}