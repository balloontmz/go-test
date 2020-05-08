/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 *
 * https://leetcode.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (38.92%)
 * Likes:    1202
 * Dislikes: 100
 * Total Accepted:    76.2K
 * Total Submissions: 195.7K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * Given a matrix consists of 0 and 1, find the distance of the nearest 0 for
 * each cell.
 * 
 * The distance between two adjacent cells is 1.
 * 
 * 
 * 
 * Example 1: 
 * 
 * 
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[0,0,0]]
 * 
 * Output:
 * [[0,0,0],
 * [0,1,0],
 * [0,0,0]]
 * 
 * 
 * Example 2: 
 * 
 * 
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,1,1]]
 * 
 * Output:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,2,1]]
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The number of elements of the given matrix will not exceed 10,000.
 * There are at least one 0 in the given matrix.
 * The cells are adjacent in only four directions: up, down, left and right.
 * 
 * 
 */

// @lc code=start
//bfs
// 48/48 cases passed (76 ms)
// Your runtime beats 43.21 % of golang submissions
// Your memory usage beats 100 % of golang submissions (8.1 MB)
const INT_MAX = int(^uint(0) >> 1)

var directions = [][]int{
	[]int{-1, 0},
	[]int{1, 0},
	[]int{0, -1},
	[]int{0, 1},
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	var queue [][]int
	//初始化数组,将所有 0 的下标存入队列
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = INT_MAX
			} else {
				queue = append(queue, []int{i, j}) // 将 0 的下标存入队列
			}
		}
	}

	for len(queue) != 0 {
		curValue := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			//没有越界
			newX := curValue[0] + direction[0]
			newY := curValue[1] + direction[1]
			if newX >= 0 && newX < len(matrix) && newY >=0 && newY < len(matrix[0]) && matrix[newX][newY] > matrix[curValue[0]][curValue[1]] {
				matrix[newX][newY] = matrix[curValue[0]][curValue[1]] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	return matrix
}

// @lc code=end

