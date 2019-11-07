/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
 *
 * https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/description/
 *
 * algorithms
 * Easy (43.07%)
 * Likes:    450
 * Dislikes: 669
 * Total Accepted:    59.1K
 * Total Submissions: 137.1K
 * Testcase Example:  '[2,2,5,null,null,5,7]'
 *
 * Given a non-empty special binary tree consisting of nodes with the
 * non-negative value, where each node in this tree has exactly two or zero
 * sub-node. If the node has two sub-nodes, then this node's value is the
 * smaller value among its two sub-nodes. More formally, the property root.val
 * = min(root.left.val, root.right.val) always holds.
 *
 * Given such a binary tree, you need to output the second minimum value in the
 * set made of all the nodes' value in the whole tree.
 *
 * If no such second minimum value exists, output -1 instead.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   5
 * ⁠    / \
 * ⁠   5   7
 *
 * Output: 5
 * Explanation: The smallest value is 2, the second smallest value is 5.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   2
 *
 * Output: -1
 * Explanation: The smallest value is 2, but there isn't any second smallest
 * value.
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
//dfs 查找二叉树中第二小的元素
func findSecondMinimumValue(root *TreeNode) int {
	// 35/35 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	if root == nil || root.Left == nil {
		return -1
	}
	leftVal := root.Left.Val
	rightVal := root.Right.Val
	if leftVal == root.Val {
		leftVal = findSecondMinimumValue(root.Left)
	}
	if rightVal == root.Val {
		rightVal = findSecondMinimumValue(root.Right)
	}
	if leftVal != -1 && rightVal != -1 {
		return min(leftVal, rightVal)
	}
	if leftVal != -1 {
		return leftVal
	}
	return rightVal
}

//bfs查找二叉树中第二小的元素 -- 切片实现，应该实现得不太好。性能还行
// var res = -1
// var bfsArr = []*TreeNode{}

// func findSecondMinimumValue(root *TreeNode) int {
// 	// 35/35 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (2 MB)
// 	var val = root.Val
// 	bfsArr = append(bfsArr, root)
// 	defer func() {
// 		res = -1
// 		bfsArr = []*TreeNode{}
// 	}()

// 	for {
// 		bfsEle := bfsArr[0]
// 		bfsArr = bfsArr[1:]
// 		bfs(bfsEle, val)
// 		if len(bfsArr) == 0 {
// 			break
// 		}
// 	}
// 	return res
// }

// func bfs(root *TreeNode, val int) {
// 	if root.Val != val {
// 		if res != -1 {
// 			res = min(res, root.Val)
// 		} else {
// 			res = root.Val
// 		}
// 		return
// 	}
// 	if root.Left == nil {
// 		return
// 	}
// 	bfsArr = append(bfsArr, root.Left)
// 	bfsArr = append(bfsArr, root.Right)
// 	return
// }

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

