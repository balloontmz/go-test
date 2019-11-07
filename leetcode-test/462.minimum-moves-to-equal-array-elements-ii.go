import "sort"

/*
 * @lc app=leetcode id=462 lang=golang
 *
 * [462] Minimum Moves to Equal Array Elements II
 *
 * https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/description/
 *
 * algorithms
 * Medium (52.85%)
 * Likes:    397
 * Dislikes: 35
 * Total Accepted:    39.2K
 * Total Submissions: 74.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty integer array, find the minimum number of moves required
 * to make all array elements equal, where a move is incrementing a selected
 * element by 1 or decrementing a selected element by 1.
 *
 * You may assume the array's length is at most 10,000.
 *
 * Example:
 *
 * Input:
 * [1,2,3]
 *
 * Output:
 * 2
 *
 * Explanation:
 * Only two moves are needed (remember each move increments or decrements one
 * element):
 *
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 */

// @lc code=start
//先排序的计算最小移动步数
func minMoves2(nums []int) int {
	// 29/29 cases passed (16 ms)
	// Your runtime beats 26.19 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.5 MB)
	sort.Ints(nums)
	var l = 0
	var h = len(nums) - 1
	var move = 0
	for l <= h {
		move += nums[h] - nums[l]
		l++
		h--
	}
	return move
}

// @lc code=end

