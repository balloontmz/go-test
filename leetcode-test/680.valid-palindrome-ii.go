/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (34.27%)
 * Likes:    739
 * Dislikes: 46
 * Total Accepted:    75.5K
 * Total Submissions: 220.2K
 * Testcase Example:  '"aba"'
 *
 * 
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 * 
 * 
 * Example 1:
 * 
 * Input: "aba"
 * Output: True
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 * 
 * 
 * 
 * Note:
 * 
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 * 
 * 
 */
func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	// 双指针同时缩小，如果能遍历完，代表该字符串是回文。
	for i < j {
		// 如果进入判断里面，则跳过一个索引继续进行回文判断
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i + 1, j)
		}
		i++
		j--
	}
	return true
}

// 判断 i,j 往内是否为回文。
func isPalindrome(s string, i, j int) bool {
	for i < j{
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

