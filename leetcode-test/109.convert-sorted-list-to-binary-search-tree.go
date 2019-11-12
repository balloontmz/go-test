/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (43.22%)
 * Likes:    1326
 * Dislikes: 75
 * Total Accepted:    200.6K
 * Total Submissions: 463.4K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given a singly linked list where elements are sorted in ascending order,
 * convert it to a height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted linked list: [-10,-3,0,5,9],
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var index *ListNode

//也可以采用中序遍历的方式进行构造，此方法无需中断链表！！！
//为什么此种方法反而运行效率更低？？？
func sortedListToBST(head *ListNode) *TreeNode {
	// 32/32 cases passed (488 ms)
	// Your runtime beats 12.73 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (273.4 MB)
	if head == nil {
		return nil
	}
	index = head
	var num = 0
	for head != nil {
		num++
		head = head.Next
	}
	return helper(0, num-1)
}

func helper(s, e int) *TreeNode {
	if s > e { // 终止条件
		return nil
	}
	var mid = (s + e + 1) / 2 // 加1左子树优先，不加右子树优先
	var left = helper(s, mid-1)
	var res = &TreeNode{
		Val: index.Val,
	}
	index = index.Next
	var right = helper(mid+1, e)
	res.Left = left
	res.Right = right
	return res
}

//从有序链表构建平衡二叉树 -- 此处采用的是中断链表的方式
//也可以采用中序遍历的方式进行构造，此方法无需中断链表！！！
// func sortedListToBST(head *ListNode) *TreeNode {
// 	// 32/32 cases passed (416 ms)
// 	// Your runtime beats 92.73 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (341.7 MB)
// 	if head == nil {
// 		return nil
// 	}
// 	if head.Next == nil {
// 		return &TreeNode{
// 			Val: head.Val,
// 		}
// 	}
// 	var preMid = preMidF(head)
// 	var mid = preMid.Next
// 	preMid.Next = nil // 中断链表
// 	var res = &TreeNode{
// 		Val: mid.Val,
// 	}
// 	res.Left = sortedListToBST(head)
// 	res.Right = sortedListToBST(mid.Next)
// 	return res
// }

// func preMidF(head *ListNode) *ListNode {
// 	// 边界条件非常重要，要注意！！！
// 	var fast, slow = head.Next, head
// 	var pre = head //初始条件！！！
// 	for fast != nil && fast.Next != nil {
// 		pre = slow
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	return pre
// }

// @lc code=end

