/*
 * @lc app=leetcode id=769 lang=golang
 *
 * [769] Max Chunks To Make Sorted
 *
 * https://leetcode.com/problems/max-chunks-to-make-sorted/description/
 *
 * algorithms
 * Medium (52.96%)
 * Likes:    573
 * Dislikes: 93
 * Total Accepted:    29K
 * Total Submissions: 54.7K
 * Testcase Example:  '[4,3,2,1,0]'
 *
 * Given an array arr that is a permutation of [0, 1, ..., arr.length - 1], we
 * split the array into some number of "chunks" (partitions), and individually
 * sort each chunk.  After concatenating them, the result equals the sorted
 * array.
 *
 * What is the most number of chunks we could have made?
 *
 * Example 1:
 *
 *
 * Input: arr = [4,3,2,1,0]
 * Output: 1
 * Explanation:
 * Splitting into two or more chunks will not return the required result.
 * For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1,
 * 2], which isn't sorted.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [1,0,2,3,4]
 * Output: 4
 * Explanation:
 * We can split into two chunks, such as [1, 0], [2, 3, 4].
 * However, splitting into [1, 0], [2], [3], [4] is the highest number of
 * chunks possible.
 *
 *
 * Note:
 *
 *
 * arr will have length in range [1, 10].
 * arr[i] will be a permutation of [0, 1, ..., arr.length - 1].
 *
 *
 *
 *
 */

// @lc code=start
//分隔数组
func maxChunksToSorted(arr []int) int {
	// 52/52 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	if len(arr) == 0 {
		return 0
	}
	var right = arr[0]
	var cnt = 0
	for i := 0; i < len(arr); i++ {
		right = max(right, arr[i])
		if right == i {
			cnt++
		}
	}
	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

