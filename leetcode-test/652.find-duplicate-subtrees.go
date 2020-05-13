import "strconv"

/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 *
 * https://leetcode.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (49.50%)
 * Likes:    1146
 * Dislikes: 192
 * Total Accepted:    57.8K
 * Total Submissions: 116.6K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * Given a binary tree, return all duplicate subtrees. For each kind of
 * duplicate subtrees, you only need to return the root node of any one of
 * them.
 * 
 * Two trees are duplicate if they have the same structure with same node
 * values.
 * 
 * Example 1: 
 * 
 * 
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   2   4
 * ⁠      /
 * ⁠     4
 * 
 * 
 * The following are two duplicate subtrees:
 * 
 * 
 * ⁠     2
 * ⁠    /
 * ⁠   4
 * 
 * 
 * and
 * 
 * 
 * ⁠   4
 * 
 * Therefore, you need to return above trees' root in the form of a list.
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
//查找重复二叉子树
//  168/168 cases passed (16 ms)
//  Your runtime beats 82.93 % of golang submissions
//  Your memory usage beats 100 % of golang submissions (11.5 MB)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res = &[]*TreeNode{}
	var M = make(map[string]int)
	helper(root, res, M)
	return *res
}

func helper(root *TreeNode, res *[]*TreeNode, M map[string]int) string {
	if root == nil {
		return "#"
	}
	var str = strconv.Itoa(root.Val) + "," + helper(root.Left, res, M) + "," + helper(root.Right, res, M)
	if M[str] == 1 {
		// fmt.Print(str, "\n")
		*res = append(*res, root)
	}
	M[str]++
	return str
}

// [-64,12,18,-4,-53,null,76,null,-51,null,null,-93,3,null,-31,47,null,3,53,-81,33,4,null,-51,-44,-60,11,null,null,null,null,78,null,-35,-64,26,-81,-31,27,60,74,null,null,8,-38,47,12,-24,null,-59,-49,-11,-51,67,null,null,null,null,null,null,null,-67,null,-37,-19,10,-55,72,null,null,null,-70,17,-4,null,null,null,null,null,null,null,3,80,44,-88,-91,null,48,-90,-30,null,null,90,-34,37,null,null,73,-38,-31,-85,-31,-96,null,null,-18,67,34,72,null,-17,-77,null,56,-65,-88,-53,null,null,null,-33,86,null,81,-42,null,null,98,-40,70,-26,24,null,null,null,null,92,72,-27,null,null,null,null,null,null,-67,null,null,null,null,null,null,null,-54,-66,-36,null,-72,null,null,43,null,null,null,-92,-1,-98,null,null,null,null,null,null,null,39,-84,null,null,null,null,null,null,null,null,null,null,null,null,null,-93,null,null,null,98]

// @lc code=end

