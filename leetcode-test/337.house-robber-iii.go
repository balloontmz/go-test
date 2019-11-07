/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 *
 * https://leetcode.com/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (49.15%)
 * Likes:    1864
 * Dislikes: 38
 * Total Accepted:    121.9K
 * Total Submissions: 247.9K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * The thief has found himself a new place for his thievery again. There is
 * only one entrance to this area, called the "root." Besides the root, each
 * house has one and only one parent house. After a tour, the smart thief
 * realized that "all houses in this place forms a binary tree". It will
 * automatically contact the police if two directly-linked houses were broken
 * into on the same night.
 *
 * Determine the maximum amount of money the thief can rob tonight without
 * alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3,null,3,null,1]
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \   \
 * ⁠    3   1
 *
 * Output: 7
 * Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
 *
 * Example 2:
 *
 *
 * Input: [3,4,5,1,3,null,1]
 *
 * 3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \   \
 * ⁠1   3   1
 *
 * Output: 9
 * Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
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
func rob(root *TreeNode) int {
	// var m = map[*TreeNode]int{}
	// return dfs(root, m)
	res := dfs(root)
	return max(res[0], res[1])
}

//采用2元素数组作为辅助空间，时间性能更优--毕竟少计算了一层递归！！！，空间上 多用了n个数组
func dfs(root *TreeNode) [2]int {
	// 124/124 cases passed (4 ms)
	// Your runtime beats 98.44 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.2 MB)
	if root == nil {
		return [2]int{}
	}
	var res [2]int
	left := dfs(root.Left)
	right := dfs(root.Right)
	res[0] = max(left[0], left[1]) + max(right[0], right[1]) // 分别取左右节点中包含与不包含的最大值！！！
	res[1] = root.Val + left[0] + right[0]
	return res
}

//采用 map 保存状态的回溯算法，以空间换取时间。保存每个已计算的节点所能获取的最大价值
//实际上进行了2层递归
// func dfs(root *TreeNode, m map[*TreeNode]int) int {
// 	// 124/124 cases passed (8 ms)
// 	// Your runtime beats 67.19 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (5.7 MB)
// 	if root == nil {
// 		return 0
// 	}
// 	if m[root] != 0 {
// 		return m[root]
// 	}
// 	//计算当前遍历节点所能获得的最大价值
// 	var v int
// 	if root.Left != nil {
// 		v += dfs(root.Left.Left, m) + dfs(root.Left.Right, m)
// 	}
// 	if root.Right != nil {
// 		v += dfs(root.Right.Left, m) + dfs(root.Right.Right, m)
// 	}
// 	v = max(v+root.Val, dfs(root.Left, m)+dfs(root.Right, m))
// 	m[root] = v
// 	return v
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end