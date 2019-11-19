/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (39.70%)
 * Likes:    2629
 * Dislikes: 41
 * Total Accepted:    204.7K
 * Total Submissions: 511.9K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two words word1 and word2, find the minimum number of operations
 * required to convert word1 to word2.
 *
 * You have the following 3 operations permitted on a word:
 *
 *
 * Insert a character
 * Delete a character
 * Replace a character
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 */

// @lc code=start
//修改一个字符串成为另一个字符串，使得修改次数最少
func minDistance(word1 string, word2 string) int {
	// 1146/1146 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.6 MB)
	var m, n = len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	//初始化dp数组
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化一个字符串为0的情况
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	//空字符串的情况已经初始化了
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] { // 如果最后一个标位相同，则可以同时去掉采用减掉该字符后的解
				dp[i][j] = dp[i-1][j-1]
			} else {
				//分别为替换一个字符、删除一个字符、添加一个字符的情况  m->n
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

