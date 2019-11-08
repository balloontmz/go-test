/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (53.38%)
 * Likes:    1528
 * Dislikes: 157
 * Total Accepted:    313.3K
 * Total Submissions: 586K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
//用数组构建平衡查找二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	// 32/32 cases passed (120 ms)
	// Your runtime beats 10.71 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (205.6 MB)
	return toBST(nums, 0, len(nums)-1)
}

func toBST(nums []int, s, e int) *TreeNode {
	if s > e {
		return nil
	}
	var root = &TreeNode{}
	root.Left = toBST(nums, s, (s+e)/2-1)
	root.Right = toBST(nums, (s+e)/2+1, e)
	root.Val = nums[(s+e)/2]
	return root
}

// @lc code=end

