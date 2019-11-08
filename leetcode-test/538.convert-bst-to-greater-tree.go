/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 *
 * https://leetcode.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (52.65%)
 * Likes:    1518
 * Dislikes: 98
 * Total Accepted:    97K
 * Total Submissions: 184.2K
 * Testcase Example:  '[5,2,13]'
 *
 * Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 * every key of the original BST is changed to the original key plus sum of all
 * keys greater than the original key in BST.
 *
 *
 * Example:
 *
 * Input: The root of a Binary Search Tree like this:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * Output: The root of a Greater Tree like this:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
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
//反转的中序遍历
var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	// 212/212 cases passed (224 ms)
	// Your runtime beats 16.24 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (272.9 MB)
	inOrder(root)
	defer func() {
		sum = 0
	}()
	return root
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Right)
	sum += root.Val
	root.Val = sum
	inOrder(root.Left)
	return
}

// @lc code=end