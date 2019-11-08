/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
 *
 * https://leetcode.com/problems/trim-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (61.20%)
 * Likes:    1324
 * Dislikes: 142
 * Total Accepted:    78.9K
 * Total Submissions: 129K
 * Testcase Example:  '[1,0,2]\n1\n2'
 *
 *
 * Given a binary search tree and the lowest and highest boundaries as L and R,
 * trim the tree so that all its elements lies in [L, R] (R >= L). You might
 * need to change the root of the tree, so the result should return the new
 * root of the trimmed binary search tree.
 *
 *
 * Example 1:
 *
 * Input:
 * ⁠   1
 * ⁠  / \
 * ⁠ 0   2
 *
 * ⁠ L = 1
 * ⁠ R = 2
 *
 * Output:
 * ⁠   1
 * ⁠     \
 * ⁠      2
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 0   4
 * ⁠  \
 * ⁠   2
 * ⁠  /
 * ⁠ 1
 *
 * ⁠ L = 1
 * ⁠ R = 3
 *
 * Output:
 * ⁠     3
 * ⁠    /
 * ⁠  2
 * ⁠ /
 * ⁠1
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
//修剪二叉树
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	// 77/77 cases passed (8 ms)
	// Your runtime beats 97.14 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

// @lc code=end

