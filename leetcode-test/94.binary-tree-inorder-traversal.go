/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (59.23%)
 * Likes:    2115
 * Dislikes: 95
 * Total Accepted:    561.3K
 * Total Submissions: 946.6K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
//二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	// 68/68 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	var stackArr []*TreeNode
	var ret []int
	for root != nil || len(stackArr) != 0 {
		for root != nil {
			stackArr = append(stackArr, root)
			root = root.Left
		}
		if len(stackArr) != 0 {
			stackEle := stackArr[len(stackArr)-1]
			stackArr = stackArr[:len(stackArr)-1]

			ret = append(ret, stackEle.Val)
			root = stackEle.Right
		}
	}
	return ret
}

// @lc code=end

