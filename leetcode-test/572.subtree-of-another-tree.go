/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 *
 * https://leetcode.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (42.70%)
 * Likes:    1507
 * Dislikes: 63
 * Total Accepted:    137.3K
 * Total Submissions: 321.3K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 *
 * Given two non-empty binary trees s and t, check whether tree t has exactly
 * the same structure and node values with a subtree of s. A subtree of s is a
 * tree consists of a node in s and all of this node's descendants. The tree s
 * could also be considered as a subtree of itself.
 *
 *
 * Example 1:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return true, because t has the same structure and node values with a subtree
 * of s.
 *
 *
 * Example 2:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 * ⁠   /
 * ⁠  0
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return false.
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
//判断一棵树是否是另一棵树的子树
var Flag = false

func isSubtree(s *TreeNode, t *TreeNode) bool {
	// 176/176 cases passed (20 ms)
	// Your runtime beats 66.67 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.2 MB)
	searchRoot(s, t)
	defer func() {
		Flag = false
	}()
	return Flag
}

func searchRoot(s *TreeNode, t *TreeNode) {
	if s == nil || t == nil {
		return
	}
	if s.Val == t.Val {
		if isEqual(s, t) {
			Flag = true
		}
	}
	searchRoot(s.Left, t)
	searchRoot(s.Right, t)
}

func isEqual(s, t *TreeNode) bool {
	if s == nil || t == nil {
		if s == t {
			return true
		}
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}

// @lc code=end

