/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 *
 * https://leetcode.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (36.89%)
 * Likes:    1355
 * Dislikes: 95
 * Total Accepted:    188.5K
 * Total Submissions: 509.2K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, return
 * the ordering of courses you should take to finish all courses.
 *
 * There may be multiple correct orders, you just need to return one of them.
 * If it is impossible to finish all courses, return an empty array.
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished
 * course 0. So the correct course order is [0,1] .
 *
 * Example 2:
 *
 *
 * Input: 4, [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,1,2,3] or [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both
 * ⁠            courses 1 and 2. Both courses 1 and 2 should be taken after you
 * finished course 0.
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3] .
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
//排序课程使完成所有课程
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 44/44 cases passed (24 ms)
	// Your runtime beats 23.89 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (7.5 MB)
	var graphic = make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graphic[i] = make([]int, 0)
	}
	// 赋值前置课程数组
	for _, pre := range prerequisites {
		graphic[pre[0]] = append(graphic[pre[0]], pre[1])
	}

	//定义某个值不构成环的全局布尔数组和某个值构成环的局部数组
	var globalMarked = make([]bool, numCourses)
	var localMarked = make([]bool, numCourses)
	var postOrder = &([]int{}) // 取地址，因为需要改变该数组
	for i := 0; i < numCourses; i++ {
		//如果存在环，那么就没有构成合法课程安排的排序
		if hasCycle(globalMarked, localMarked, graphic, i, postOrder) {
			return []int{}
		}
	}
	return *postOrder
}

func hasCycle(globalMarked, localMarked []bool, graphic [][]int, curNode int, postOrder *[]int) bool {
	if localMarked[curNode] {
		return true
	}
	if globalMarked[curNode] {
		return false
	}
	localMarked[curNode] = true
	globalMarked[curNode] = true
	for _, nextNode := range graphic[curNode] {
		if hasCycle(globalMarked, localMarked, graphic, nextNode, postOrder) {
			return true
		}
	}
	localMarked[curNode] = false
	*postOrder = append(*postOrder, curNode) // 深度优先遍历，那么最后遍历的就是最后没有前置课程的课
	return false
}

// @lc code=end

