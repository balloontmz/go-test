/*
 * @lc app=leetcode id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 *
 * https://leetcode.com/problems/is-graph-bipartite/description/
 *
 * algorithms
 * Medium (45.17%)
 * Likes:    860
 * Dislikes: 114
 * Total Accepted:    68.3K
 * Total Submissions: 150.9K
 * Testcase Example:  '[[1,3],[0,2],[1,3],[0,2]]'
 *
 * Given an undirected graph, return true if and only if it is bipartite.
 *
 * Recall that a graph is bipartite if we can split it's set of nodes into two
 * independent subsets A and B such that every edge in the graph has one node
 * in A and another node in B.
 *
 * The graph is given in the following form: graph[i] is a list of indexes j
 * for which the edge between nodes i and j exists.  Each node is an integer
 * between 0 and graph.length - 1.  There are no self edges or parallel edges:
 * graph[i] does not contain i, and it doesn't contain any element twice.
 *
 *
 * Example 1:
 * Input: [[1,3], [0,2], [1,3], [0,2]]
 * Output: true
 * Explanation:
 * The graph looks like this:
 * 0----1
 * |    |
 * |    |
 * 3----2
 * We can divide the vertices into two groups: {0, 2} and {1, 3}.
 *
 *
 *
 * Example 2:
 * Input: [[1,2,3], [0,2], [0,1,3], [0,2]]
 * Output: false
 * Explanation:
 * The graph looks like this:
 * 0----1
 * | \  |
 * |  \ |
 * 3----2
 * We cannot find a way to divide the set of nodes into two independent
 * subsets.
 *
 *
 *
 *
 * Note:
 *
 *
 * graph will have length in range [1, 100].
 * graph[i] will contain integers in range [0, graph.length - 1].
 * graph[i] will not contain i or duplicate values.
 * The graph is undirected: if any element j is in graph[i], then i will be in
 * graph[j].
 *
 *
 */

// @lc code=start
//判断是否为二分图
func isBipartite(graph [][]int) bool {
	// 78/78 cases passed (20 ms)
	// Your runtime beats 98.04 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	var n = len(graph)
	var colorArr = make([]int, n)
	//将颜色数组标记为未染色
	for i := 0; i < n; i++ {
		colorArr[i] = -1
	}
	for i := 0; i < n; i++ {
		if colorArr[i] == -1 && !dyeing(i, 0, colorArr, graph) {
			return false
		}
	}
	return true
}

//进行染色
func dyeing(curNode, curColor int, colorArr []int, graph [][]int) bool {
	//如果当前节点已染色，则判断是否与即将染色的颜色匹配
	if colorArr[curNode] != -1 {
		return colorArr[curNode] == curColor
	}
	colorArr[curNode] = curColor
	for _, nextNode := range graph[curNode] {
		if !dyeing(nextNode, 1-curColor, colorArr, graph) {
			return false
		}
	}
	return true
}

// @lc code=end

