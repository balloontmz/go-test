/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (53.60%)
 * Likes:    1052
 * Dislikes: 119
 * Total Accepted:    108.5K
 * Total Submissions: 202.2K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given a Binary Search Tree and a target number, return true if there exist
 * two elements in the BST such that their sum is equal to the given target.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 9
 *
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 28
 *
 * Output: False
 *
 *
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
//判断一棵树中是否存在两个节点和为给定数 -- 先中序遍历，然后采用双指针。
func findTarget(root *TreeNode, k int) bool {
	// 421/421 cases passed (24 ms)
	// Your runtime beats 84.17 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (7.2 MB)
	var treeList = inOrder(root)
	for i, j := 0, len(treeList)-1; i < j; {
		sum := treeList[i] + treeList[j]
		if sum == k {
			return true
		}
		if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode // 此处不可加入根节点，否则会存在两个根节点！！！
	var res = []int{}
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			var node = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
			res = append(res, node.Val)
			root = node.Right
		}
	}

	return res
}

// @lc code=end

