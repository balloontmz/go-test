/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 *
 * https://leetcode.com/problems/reshape-the-matrix/description/
 *
 * algorithms
 * Easy (59.49%)
 * Likes:    643
 * Dislikes: 93
 * Total Accepted:    87.2K
 * Total Submissions: 146.6K
 * Testcase Example:  '[[1,2],[3,4]]\n1\n4'
 *
 * In MATLAB, there is a very useful function called 'reshape', which can
 * reshape a matrix into a new one with different size but keep its original
 * data.
 *
 *
 *
 * You're given a matrix represented by a two-dimensional array, and two
 * positive integers r and c representing the row number and column number of
 * the wanted reshaped matrix, respectively.
 *
 * ⁠The reshaped matrix need to be filled with all the elements of the original
 * matrix in the same row-traversing order as they were.
 *
 *
 *
 * If the 'reshape' operation with given parameters is possible and legal,
 * output the new reshaped matrix; Otherwise, output the original matrix.
 *
 *
 * Example 1:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 1, c = 4
 * Output:
 * [[1,2,3,4]]
 * Explanation:The row-traversing of nums is [1,2,3,4]. The new reshaped matrix
 * is a 1 * 4 matrix, fill it row by row by using the previous list.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 2, c = 4
 * Output:
 * [[1,2],
 * ⁠[3,4]]
 * Explanation:There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So
 * output the original matrix.
 *
 *
 *
 * Note:
 *
 * The height and width of the given matrix is in range [1, 100].
 * The given r and c are all positive.
 *
 *
 */

// @lc code=start
// 改变矩阵维度
func matrixReshape(nums [][]int, r int, c int) [][]int {
	// 56/56 cases passed (64 ms)
	// Your runtime beats 37.7 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (135.2 MB)
	var m, n = len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	var res = make([][]int, r)
	var index int
	for k, _ := range res {
		res[k] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i][j] = nums[index/n][index%n]
			index++
		}
	}
	return res
}

// @lc code=end

