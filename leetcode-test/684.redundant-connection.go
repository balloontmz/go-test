/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 *
 * https://leetcode.com/problems/redundant-connection/description/
 *
 * algorithms
 * Medium (54.11%)
 * Likes:    871
 * Dislikes: 207
 * Total Accepted:    67.5K
 * Total Submissions: 124.2K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 *
 * In this problem, a tree is an undirected graph that is connected and has no
 * cycles.
 *
 * The given input is a graph that started as a tree with N nodes (with
 * distinct values 1, 2, ..., N), with one additional edge added.  The added
 * edge has two different vertices chosen from 1 to N, and was not an edge that
 * already existed.
 *
 * The resulting graph is given as a 2D-array of edges.  Each element of edges
 * is a pair [u, v] with u < v, that represents an undirected edge connecting
 * nodes u and v.
 *
 * Return an edge that can be removed so that the resulting graph is a tree of
 * N nodes.  If there are multiple answers, return the answer that occurs last
 * in the given 2D-array.  The answer edge [u, v] should be in the same format,
 * with u < v.
 * Example 1:
 *
 * Input: [[1,2], [1,3], [2,3]]
 * Output: [2,3]
 * Explanation: The given undirected graph will be like this:
 * ⁠ 1
 * ⁠/ \
 * 2 - 3
 *
 *
 * Example 2:
 *
 * Input: [[1,2], [2,3], [3,4], [1,4], [1,5]]
 * Output: [1,4]
 * Explanation: The given undirected graph will be like this:
 * 5 - 1 - 2
 * ⁠   |   |
 * ⁠   4 - 3
 *
 *
 * Note:
 * The size of the input 2D-array will be between 3 and 1000.
 * Every integer represented in the 2D-array will be between 1 and N, where N
 * is the size of the input array.
 *
 *
 *
 *
 *
 * Update (2017-09-26):
 * We have overhauled the problem description + test cases and specified
 * clearly the graph is an undirected graph. For the directed graph follow up
 * please see Redundant Connection II). We apologize for any inconvenience
 * caused.
 *
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	//[dfs，bfs和并查集](https://www.cnblogs.com/grandyang/p/7628977.html)
	//深度优先遍历查找冗余连接
	// 39/39 cases passed (4 ms)
	// Your runtime beats 94.34 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.1 MB)
	var m = make(map[int]([]int), 0) // 此map用处？
	for _, edge := range edges {
		//深度优先遍历
		if hasCycle(edge[0], edge[1], m, -1) {
			return edge // 如果存在环，那么当前边是环的最后一条边
		}
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}
	return nil // 代表没环存在
}

func hasCycle(curNode, target int, m map[int]([]int), pre int) bool {
	// 如果当前节点的目标数组中存在目标，那么存在环。

	if targetArr, ok := m[curNode]; ok {
		for _, t := range targetArr {
			if target == t {
				return true
			}
			if pre == t {
				continue
			}
			if hasCycle(t, target, m, curNode) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

