/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (54.15%)
 * Likes:    1496
 * Dislikes: 49
 * Total Accepted:    271.1K
 * Total Submissions: 500K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given a binary search tree, write a function kthSmallest to find the kth
 * smallest element in it.
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
 *
 * Example 1:
 *
 *
 * Input: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * Output: 1
 *
 * Example 2:
 *
 *
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * Output: 3
 *
 *
 * Follow up:
 * What if the BST is modified (insert/delete operations) often and you need to
 * find the kth smallest frequently? How would you optimize the kthSmallest
 * routine?
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
//递归查找第n小的元素 -- 性能更差
func kthSmallest(root *TreeNode, k int) int {
	// 91/91 cases passed (12 ms)
	// Your runtime beats 68.16 % of golang submissions
	// Your memory usage beats 80 % of golang submissions (5.8 MB)
	leftCount := count(root.Left)
	if leftCount == k-1 {
		return root.Val
	}
	if leftCount > k-1 {
		return kthSmallest(root.Left, k)
	}
	return kthSmallest(root.Right, k-leftCount-1)
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.Left) + count(root.Right)
}

//中序遍历查找第n小的元素
// func kthSmallest(root *TreeNode, k int) int {
// 	// 为什么循环比递归性能差！！！ 应该是切片操作的影响
// 	// 91/91 cases passed (12 ms)
// 	// Your runtime beats 68.16 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (5.8 MB)
// 	var cnt = 0
// 	var stackArr = []*TreeNode{}
// 	for len(stackArr) != 0 || root != nil {
// 		for root != nil {
// 			stackArr = append(stackArr, root)
// 			root = root.Left
// 		}
// 		root = stackArr[len(stackArr)-1]
// 		stackArr = stackArr[:len(stackArr)-1]
// 		cnt++
// 		if cnt == k {
// 			return root.Val
// 		}
// 		root = root.Right
// 	}
// 	return 0
// }

// var cnt = 0
// var val = 0

// func kthSmallest(root *TreeNode, k int) int {
// 	// 91/91 cases passed (8 ms)
// 	// Your runtime beats 98.65 % of golang submissions
// 	// Your memory usage beats 90 % of golang submissions (5.8 MB)
// 	inOrder(root, k)
// 	defer func() {
// 		cnt = 0
// 		val = 0
// 	}()
// 	return val
// }

// func inOrder(root *TreeNode, k int) {
// 	if root == nil {
// 		return
// 	}
// 	inOrder(root.Left, k)
// 	cnt++
// 	if cnt == k {
// 		val = root.Val
// 		return
// 	}
// 	inOrder(root.Right, k)
// 	return
// }

// @lc code=end

