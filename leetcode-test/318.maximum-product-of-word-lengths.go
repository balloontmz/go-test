/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 *
 * https://leetcode.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (49.54%)
 * Likes:    640
 * Dislikes: 52
 * Total Accepted:    87.2K
 * Total Submissions: 175.8K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * Given a string array words, find the maximum value of length(word[i]) *
 * length(word[j]) where the two words do not share common letters. You may
 * assume that each word will contain only lower case letters. If no such two
 * words exist, return 0.
 *
 * Example 1:
 *
 *
 * Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * Output: 16
 * Explanation: The two words can be "abcw", "xtfn".
 *
 * Example 2:
 *
 *
 * Input: ["a","ab","abc","d","cd","bcd","abcd"]
 * Output: 4
 * Explanation: The two words can be "ab", "cd".
 *
 * Example 3:
 *
 *
 * Input: ["a","aa","aaa","aaaa"]
 * Output: 0
 * Explanation: No such pair of words.
 *
 *
 */

// @lc code=start
//字符串数组最大乘积
func maxProduct(words []string) int {
	// 174/174 cases passed (20 ms)
	// Your runtime beats 38.3 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6 MB)
	var val = make([]uint32, len(words))
	for k, word := range words {
		for i := 0; i < len(word); i++ {
			//采用32位整数存储每个字符串对应的字符 -- 字符最多为26个。
			val[k] |= 1 << (word[i] - 'a')
		}
	}
	var m int
	for i := 0; i < len(words); i++ {
		for j := i; j < len(words); j++ {
			//并为0代表没有重复字符
			if val[i]&val[j] == 0 {
				m = max(m, len(words[i])*len(words[j]))
			}
		}
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

