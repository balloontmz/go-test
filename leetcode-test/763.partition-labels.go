/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 *
 * https://leetcode.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (70.48%)
 * Likes:    1114
 * Dislikes: 59
 * Total Accepted:    60K
 * Total Submissions: 83.7K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * 
 * A string S of lowercase letters is given.  We want to partition this string
 * into as many parts as possible so that each letter appears in at most one
 * part, and return a list of integers representing the size of these parts.
 * 
 * 
 * Example 1:
 * 
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it
 * splits S into less parts.
 * 
 * 
 * 
 * Note:
 * S will have length in range [1, 500].
 * S will consist of lowercase letters ('a' to 'z') only.
 * 
 */
func partitionLabels(S string) []int {
	// ✔ 116/116 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.3 MB)
	var maxMap = map[byte]int{}
	for i := 0; i < len(S); i++ {
		maxMap[S[i]] = i
	}
	var start, end int
	var res = []int{}
	for i := 0; i < len(S); i++ {
		end = max(maxMap[S[i]], end) // 取当前遍历的最大索引或之前存在的最大索引
		if end == i {  // 相等代表已经是一个完整的子序列了
			res = append(res, end - start + 1)
			start = end + 1
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
