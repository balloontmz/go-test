/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 *
 * https://leetcode.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (47.68%)
 * Likes:    1808
 * Dislikes: 111
 * Total Accepted:    173.2K
 * Total Submissions: 363.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 *
 * Given a binary tree, you need to compute the length of the diameter of the
 * tree. The diameter of a binary tree is the length of the longest path
 * between any two nodes in a tree. This path may or may not pass through the
 * root.
 *
 *
 *
 * Example:
 * Given a binary tree
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 *
 * Note:
 * The length of path between two nodes is represented by the number of edges
 * between them.
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
//计算树最大两个节点距离 -- 其实就是计算节点高度和的最大值！！！
var maxPath = 0

func diameterOfBinaryTree(root *TreeNode) int {
	// 106/106 cases passed (4 ms)
	// Your runtime beats 94.56 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.5 MB)
	maxDepth(root)
	defer func() {
		maxPath = 0
	}()
	return maxPath
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	maxPath = max(maxPath, l+r)

	return max(l, r) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

