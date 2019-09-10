/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (46.76%)
 * Likes:    1014
 * Dislikes: 77
 * Total Accepted:    245.3K
 * Total Submissions: 520.3K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * Output: ["1->2->5", "1->3"]
 * 
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	// ✔ 209/209 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.5 MB)
	var paths = &[]string{}
	if root == nil {
		return *paths
	}
	var values = &[]string{}
	backtracking(root, values, paths)
	return *paths
}

func backtracking(node *TreeNode, values, paths *[]string)  {
	if node == nil {
		return
	}
	s :=strconv.Itoa(node.Val)
	*values = append(*values, s)
	if isLeaf(node) {
		*paths = append(*paths, buildPath(values))
	} else {
		backtracking(node.Left, values, paths)
		backtracking(node.Right, values, paths)	
	}
	*values = (*values)[:len(*values)-1]
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func buildPath(values *[]string) string {
	var str string
	for i := 0; i < len(*values); i++ {
		str += (*values)[i]
		if i != len(*values) - 1 {
			str += "->"
		}
	}
	return str
}

