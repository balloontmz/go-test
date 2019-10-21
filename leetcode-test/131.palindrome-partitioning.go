/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (43.00%)
 * Likes:    1157
 * Dislikes: 48
 * Total Accepted:    184.9K
 * Total Submissions: 429.4K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 *
 */

// @lc code=start
func partition(s string) [][]string {
	var partions [][]string
	var tempPartion []string
	doPartion(s, &partions, tempPartion)
	return partions
}

func doPartion(s string, partions *[][]string, tempPartion []string) {
	if len(s) == 0 {
		*partions = append(*partions, append([]string{}, tempPartion...)) // 参数展开并合并成新数组!!!原切片只指向一个地址
		return
	}

	for i := 0; i < len(s); i++ {
		if isPalindrome(s, 0, i) {
			tempPartion = append(tempPartion, s[0:i+1])
			doPartion(s[i+1:], partions, tempPartion)
			tempPartion = tempPartion[:len(tempPartion)-1]
		}
	}
}

func isPalindrome(s string, begin, end int) bool {
	for i, j := begin, end; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

