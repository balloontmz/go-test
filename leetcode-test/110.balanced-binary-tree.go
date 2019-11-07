/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (42.00%)
 * Likes:    1541
 * Dislikes: 140
 * Total Accepted:    369.9K
 * Total Submissions: 880.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 *
 *
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
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
//判断一棵树是否是平衡树 -- 需要判断每个节点是否是平衡的！！！
var result = true

func isBalanced(root *TreeNode) bool {
	// 227/227 cases passed (4 ms)
	// Your runtime beats 99.37 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	maxDepth(root)
	tempResult := result
	result = true
	return tempResult
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l-r < -1 || l-r > 1 {
		result = false
	}

	return max(l, r) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

