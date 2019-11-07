/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 *
 * https://leetcode.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Easy (34.66%)
 * Likes:    1188
 * Dislikes: 308
 * Total Accepted:    70.8K
 * Total Submissions: 204.3K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * Given a binary tree, find the length of the longest path where each node in
 * the path has the same value. This path may or may not pass through the
 * root.
 *
 * The length of path between two nodes is represented by the number of edges
 * between them.
 *
 *
 *
 * Example 1:
 *
 * Input:
 *
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         1   1   5
 *
 *
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 * Input:
 *
 *
 * ⁠             1
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         4   4   5
 *
 *
 * Output: 2
 *
 *
 *
 * Note: The given binary tree has not more than 10000 nodes. The height of the
 * tree is not more than 1000.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//最长同值路径
var sum = 0

func longestUnivaluePath(root *TreeNode) int {
	// 71/71 cases passed (72 ms)
	// Your runtime beats 89.53 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.9 MB)
	dfs(root)
	defer func() {
		sum = 0
	}()
	return sum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if root.Left != nil && root.Left.Val == root.Val {
		left = left + 1
	} else {
		left = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right = right + 1
	} else {
		right = 0
	}
	//返回值和计算值不一致 -- 最终被采用的为计算值
	sum = max(sum, left+right)
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

