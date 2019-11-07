/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Easy (44.10%)
 * Likes:    2344
 * Dislikes: 148
 * Total Accepted:    132.1K
 * Total Submissions: 299.4K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
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
//计算二叉树中路径和为给定和的个数
//重点是开始点不一定是root节点，所以需要递归指定开始点
//由于计算链是连续的，所以可以再采用一个递归进行计算！！！
var count = 0

func pathSum(root *TreeNode, sum int) int {
	// 126/126 cases passed (16 ms)
	// Your runtime beats 68.39 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.3 MB)
	if root == nil {
		return 0
	}
	//此处应该递归！！！开始的节点可以是任意的！！！
	accessCount(root, sum)
	// countNum(root, sum)
	// countNum(root.Left, sum)
	// countNum(root.Right, sum)
	defer func() {
		count = 0
	}()
	return count
}

func countNum(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		count += 1
	}
	// 可能包括父节点、也可能不包括
	countNum(root.Left, sum-root.Val)
	countNum(root.Right, sum-root.Val)
	return
}

func accessCount(root *TreeNode, sum int) {
	countNum(root, sum)
	if root.Left != nil {
		accessCount(root.Left, sum)
	}
	if root.Right != nil {
		accessCount(root.Right, sum)
	}
}

// @lc code=end

