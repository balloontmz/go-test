/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (36.50%)
 * Likes:    1388
 * Dislikes: 119
 * Total Accepted:    148.1K
 * Total Submissions: 404.8K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 * 
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
func generateTrees(n int) []*TreeNode {
    if n < 1 {
		return []*TreeNode{}
	}
	return generateSubtrees(1, n)
}

func generateSubtrees(l, r int) []*TreeNode {
	var res = []*TreeNode{}
	// fmt.Print("测试进入", l, "-----", r)
	// 最小元素
	if l > r {
		// 此处nil 和 &TreeNode{} 的区别为 TreeNode 会自带空值的 left 和 right
		res = append(res, nil)
		// fmt.Print("测试返回", res)
		return res
	}

	// 此处边界条件 i < r 会导致无法计算出结果--i=r 的情况没有被覆盖
	for i := l; i <= r; i++ {
		left := generateSubtrees(l, i-1)
		right := generateSubtrees(i+1, r)
		// fmt.Print("left", left, "right", right)
		for _, rv := range right {
			for _, lv := range left {
				root := &TreeNode{Val: i}
				root.Left = lv
				root.Right = rv
				res = append(res, root)
				// fmt.Println("root", root, "left", lv, "right", rv)
			}
		}
	}
	// fmt.Print("测试返回", res)
	return res
}

