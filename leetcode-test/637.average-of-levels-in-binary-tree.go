/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 *
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (60.12%)
 * Likes:    916
 * Dislikes: 132
 * Total Accepted:    94.1K
 * Total Submissions: 156.5K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * Given a non-empty binary tree, return the average value of the nodes on each
 * level in the form of an array.
 *
 * Example 1:
 *
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
 * level 2 is 11. Hence return [3, 14.5, 11].
 *
 *
 *
 * Note:
 *
 * The range of node's value is in the range of 32-bit signed integer.
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
//bfs 计算二叉树每层的平均值
func averageOfLevels(root *TreeNode) []float64 {
	// 65/65 cases passed (12 ms)
	// Your runtime beats 96.67 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (7.7 MB)
	var count = 0              // 层计数，到0代表进入了下一层
	var level = -1             // 当前所在的层，0开始
	var bfsArr = []*TreeNode{} // 遍历队列
	var ret = []float64{}      // 结果数组
	var temp = 0               // 计算平均值的临时变量
	var div = 0                // 计算平均值的临时除数
	bfsArr = append(bfsArr, root)
	for len(bfsArr) != 0 {
		if count == 0 {
			count = len(bfsArr)
			level++ // 到了下一层
			temp = 0
			div = len(bfsArr)
		}
		count--
		bfsEle := bfsArr[0]
		bfsArr = bfsArr[1:]
		if bfsEle.Left != nil {
			bfsArr = append(bfsArr, bfsEle.Left)
		}
		if bfsEle.Right != nil {
			bfsArr = append(bfsArr, bfsEle.Right)
		}
		temp += bfsEle.Val
		if count == 0 {
			ret = append(ret, float64(temp)/float64(div))

		}
	}
	return ret
}

// @lc code=end

