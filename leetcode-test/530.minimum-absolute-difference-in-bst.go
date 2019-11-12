/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (51.73%)
 * Likes:    624
 * Dislikes: 50
 * Total Accepted:    70.7K
 * Total Submissions: 136.5K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * Output:
 * 1
 *
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 *
 *
 *
 *
 * Note: There are at least two nodes in this BST.
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
var preNode *TreeNode
var m = int(^uint(0) >> 1)

func getMinimumDifference(root *TreeNode) int {
	//利用前序遍历的是有序的，所以分别计算前序遍历序列中两两相邻的节点差即可
	//遍历两两相邻，节点可能并不相邻！！！
	// 186/186 cases passed (16 ms)
	// Your runtime beats 21.95 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.6 MB)
	// inorder2(root) //性能更优
	inOrder(root)
	defer func() {
		preNode = nil
		m = int(^uint(0) >> 1)
	}()
	return m
}

func inorder2(root *TreeNode) {
	//性能更优！！！
	// 186/186 cases passed (12 ms)
	// Your runtime beats 80.49 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.6 MB)
	if root == nil {
		return
	}
	inorder2(root.Left)
	if preNode != nil {
		m = min(m, root.Val-preNode.Val)
	}
	preNode = root
	inorder2(root.Right)
	return
}

func inOrder(root *TreeNode) {
	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if preNode != nil {
			m = min(m, root.Val-preNode.Val)
		}
		preNode = root
		root = root.Right
	}
	return
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

