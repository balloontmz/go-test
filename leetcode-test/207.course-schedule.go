/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (39.79%)
 * Likes:    2432
 * Dislikes: 123
 * Total Accepted:    283.6K
 * Total Submissions: 710.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, is it
 * possible for you to finish all courses?
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 *
 * Input: 2, [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should
 * also have finished course 1. So it is impossible.
 *
 *
 * Note:
 *
 *
 * The input prerequisites is a graph represented by a list of edges, not
 * adjacency matrices. Read more about how a graph is represented.
 * You may assume that there are no duplicate edges in the input
 * prerequisites.
 *
 *
 */

// @lc code=start
//课程是否合法--图是否存在环。 -- 带标记的深度优先遍历 -- 回溯？？？
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 42/42 cases passed (12 ms)
	// Your runtime beats 84.9 % of golang submissions
	// Your memory usage beats 25 % of golang submissions (6.2 MB)
	var graphic = make([][]int, numCourses)
	//初始化前置课程数组
	for i := 0; i < numCourses; i++ {
		graphic[i] = make([]int, 0)
	}
	// 将前置课程加入数组
	for _, pre := range prerequisites {
		graphic[pre[0]] = append(graphic[pre[0]], pre[1])
	}

	var globalMarked = make([]bool, numCourses)
	var localMarked = make([]bool, numCourses)
	// 判断是否有环
	for i := 0; i < numCourses; i++ {
		if hasCycle(globalMarked, localMarked, graphic, i) {
			return false
		}
	}
	return true
}

func hasCycle(globalMarked, localMarked []bool, graphic [][]int, curNode int) bool {
	//如果当前遍历已经存在，则代表存在环
	if localMarked[curNode] {
		return true
	}
	//如果当前遍历不存在，但是全局存在了，则代表不存在环--全局的那个是没有环的路径
	if globalMarked[curNode] {
		return false
	}
	//当前遍历和全局遍历标记为已遍历
	globalMarked[curNode] = true
	localMarked[curNode] = true
	for _, nextNode := range graphic[curNode] {
		// 如果子节点存在环
		if hasCycle(globalMarked, localMarked, graphic, nextNode) {
			return true
		}
	}
	//当前节点遍历完成，不存在环。取消标记当前节点，此时全局的该节点为不存在环的节点。
	localMarked[curNode] = false
	return false
}

// @lc code=end

