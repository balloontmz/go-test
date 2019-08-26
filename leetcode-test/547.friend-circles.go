/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Friend Circles
 *
 * https://leetcode.com/problems/friend-circles/description/
 *
 * algorithms
 * Medium (54.44%)
 * Likes:    1141
 * Dislikes: 94
 * Total Accepted:    98.6K
 * Total Submissions: 180.5K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 * 
 * There are N students in a class. Some of them are friends, while some are
 * not. Their friendship is transitive in nature. For example, if A is a direct
 * friend of B, and B is a direct friend of C, then A is an indirect friend of
 * C. And we defined a friend circle is a group of students who are direct or
 * indirect friends.
 * 
 * 
 * 
 * Given a N*N matrix M representing the friend relationship between students
 * in the class. If M[i][j] = 1, then the ith and jth students are direct
 * friends with each other, otherwise not. And you have to output the total
 * number of friend circles among all the students.
 * 
 * 
 * Example 1:
 * 
 * Input: 
 * [[1,1,0],
 * ⁠[1,1,0],
 * ⁠[0,0,1]]
 * Output: 2
 * Explanation:The 0th and 1st students are direct friends, so they are in a
 * friend circle. The 2nd student himself is in a friend circle. So return
 * 2.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: 
 * [[1,1,0],
 * ⁠[1,1,1],
 * ⁠[0,1,1]]
 * Output: 1
 * Explanation:The 0th and 1st students are direct friends, the 1st and 2nd
 * students are direct friends, so the 0th and 2nd students are indirect
 * friends. All of them are in the same friend circle, so return 1.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * N is in range [1,200].
 * M[i][i] = 1 for all students.
 * If M[i][j] = 1, then M[j][i] = 1.
 * 
 * 
 */
var n int

func findCircleNum(M [][]int) int {
	// ✔ 113/113 cases passed (24 ms)
	// ✔ Your runtime beats 35.81 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (6.2 MB)
	n = len(M)
	var circleNum = 0
	var hasVisited = make([]bool, n)
	for i := 0; i < n; i++ {
		if !hasVisited[i] {
			dfs(M, i, hasVisited)
			circleNum++
		}
	}
	return circleNum
}

func dfs(M [][]int, i int, hasVisited []bool)  {
	hasVisited[i] = true
	for k := 0; k < n; k++ {
		if M[i][k] == 1 && !hasVisited[k] {
			dfs(M, k, hasVisited)
		}
	}
}

