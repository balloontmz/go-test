/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (50.89%)
 * Likes:    1180
 * Dislikes: 63
 * Total Accepted:    302.6K
 * Total Submissions: 593.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the postorder traversal of its nodes' values.
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
 * Output: [3,2,1]
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
//前序遍历变种，然后倒转结果数组
func postorderTraversal(root *TreeNode) []int {
	// 68/68 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 33.33 % of golang submissions (2.1 MB)
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
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	return reverse(ret)
}

func reverse(ret []int) []int {
	s := []int{}
	for i := len(ret) - 1; i >= 0; i-- {
		s = append(s, ret[i])
	}
	return s
}

// @lc code=end

