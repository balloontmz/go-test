/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 *
 * https://leetcode.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (60.25%)
 * Likes:    2167
 * Dislikes: 36
 * Total Accepted:    379.3K
 * Total Submissions: 629.1K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * Invert a binary tree.
 *
 * Example:
 *
 * Input:
 *
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 *
 * Output:
 *
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 *
 * Trivia:
 * This problem was inspired by this original tweet by Max Howell:
 *
 * Google: 90% of our engineers use the software you wrote (Homebrew), but you
 * can’t invert a binary tree on a whiteboard so f*** off.
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
//反转二叉树 -- 递归
func invertTree(root *TreeNode) *TreeNode {
	// 68/68 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 66.67 % of golang submissions (2.2 MB)
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	var left = root.Left
	root.Left = root.Right
	root.Right = left
	return root
}

// @lc code=end

