/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (53.04%)
 * Likes:    1006
 * Dislikes: 46
 * Total Accepted:    394.4K
 * Total Submissions: 743.2K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the preorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,2,3]
 *
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
//采用栈实现二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	// 68/68 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	var ret []int
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil { // left 后加入，先弹出
			stack = append(stack, root.Left)
		}
	}
	return ret
}

// @lc code=end

