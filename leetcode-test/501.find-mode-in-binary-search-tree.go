/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 *
 * https://leetcode.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (40.43%)
 * Likes:    650
 * Dislikes: 262
 * Total Accepted:    67.1K
 * Total Submissions: 165.8K
 * Testcase Example:  '[1,null,2,2]'
 *
 * Given a binary search tree (BST) with duplicates, find all the mode(s) (the
 * most frequently occurred element) in the given BST.
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than or equal
 * to the node's key.
 * The right subtree of a node contains only nodes with keys greater than or
 * equal to the node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * For example:
 * Given BST [1,null,2,2],
 *
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 *
 *
 * return [2].
 *
 * Note: If a tree has more than one mode, you can return them in any order.
 *
 * Follow up: Could you do that without using any extra space? (Assume that the
 * implicit stack space incurred due to recursion does not count).
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
//前序遍历统计出现次数最多的节点--可能存在多个
var curCnt = 1
var maxCnt = 1
var preNode *TreeNode
var nums = []int{}

func findMode(root *TreeNode) []int {
	// 25/25 cases passed (8 ms)
	// Your runtime beats 97.89 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	inOrder(root)
	defer func() {
		curCnt = 1
		maxCnt = 1
		preNode = nil
		nums = []int{}
	}()
	return nums
}

//前序遍历
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	if preNode != nil {
		if preNode.Val == root.Val {
			curCnt++
		} else {
			curCnt = 1
		}
	}
	if curCnt > maxCnt {
		maxCnt = curCnt
		nums = []int{root.Val} // 清空并重新赋值
	} else if curCnt == maxCnt {
		nums = append(nums, root.Val)
	}
	preNode = root

	inOrder(root.Right)
	return
}

// @lc code=end

