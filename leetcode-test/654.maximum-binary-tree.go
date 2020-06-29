/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
 *
 * https://leetcode.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (78.96%)
 * Likes:    1753
 * Dislikes: 228
 * Total Accepted:    129.1K
 * Total Submissions: 162.6K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 
 * Given an integer array with no duplicates. A maximum tree building on this
 * array is defined as follow:
 * 
 * The root is the maximum number in the array. 
 * The left subtree is the maximum tree constructed from left part subarray
 * divided by the maximum number.
 * The right subtree is the maximum tree constructed from right part subarray
 * divided by the maximum number. 
 * 
 * 
 * 
 * 
 * Construct the maximum tree by the given array and output the root node of
 * this tree.
 * 
 * 
 * Example 1:
 * 
 * Input: [3,2,1,6,0,5]
 * Output: return the tree root node representing the following tree:
 * 
 * ⁠     6
 * ⁠   /   \
 * ⁠  3     5
 * ⁠   \    / 
 * ⁠    2  0   
 * ⁠      \
 * ⁠       1
 * 
 * 
 * 
 * Note:
 * 
 * The size of the given array will be in the range [1,1000].
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
//最大值二叉树 -- 递归
// 107/107 cases passed (56 ms)
// Your runtime beats 98.97 % of golang submissions
// Your memory usage beats 89.8 % of golang submissions (8.8 MB)
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}
	//接下来是长度不为 0 的情况
	var index = findMaximumIndex(nums)
	var root = &TreeNode{
		Val: nums[index],
	}
	root.Left = constructMaximumBinaryTree(nums[0:index])
	if len(nums) > index + 1 {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return root
}

func findMaximumIndex(nums []int) int {
	var index = 0
	var num = nums[0]
	for k, v := range nums {
		if v > num {
			num = v
			index = k
		}
	}
	return index
}
// @lc code=end

