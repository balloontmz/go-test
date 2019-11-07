/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
 *
 * https://leetcode.com/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (59.64%)
 * Likes:    667
 * Dislikes: 113
 * Total Accepted:    83.9K
 * Total Submissions: 140.6K
 * Testcase Example:  '[2,1,3]'
 *
 *
 * Given a binary tree, find the leftmost value in the last row of the tree.
 *
 *
 * Example 1:
 *
 * Input:
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * Output:
 * 1
 *
 *
 *
 * ⁠ Example 2:
 *
 * Input:
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   5   6
 * ⁠      /
 * ⁠     7
 *
 * Output:
 * 7
 *
 *
 *
 * Note:
 * You may assume the tree (i.e., the given root node) is not NULL.
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
//找出二叉树左下角的元素
func findBottomLeftValue(root *TreeNode) int {
	// 74/74 cases passed (8 ms)
	// Your runtime beats 79.17 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (5.7 MB)
	var bfsArr []*TreeNode
	bfsArr = append(bfsArr, root)
	//bfs，left节点后入队列，那么左下角的元素最后入队！！！
	for len(bfsArr) != 0 {
		root = bfsArr[0]
		bfsArr = bfsArr[1:]
		if root.Right != nil {
			bfsArr = append(bfsArr, root.Right)
		}
		if root.Left != nil {
			bfsArr = append(bfsArr, root.Left)
		}
	}
	return root.Val
}

// @lc code=end

