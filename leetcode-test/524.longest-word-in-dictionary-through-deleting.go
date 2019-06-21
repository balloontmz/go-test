/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 *
 * https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/
 *
 * algorithms
 * Medium (45.83%)
 * Likes:    326
 * Dislikes: 175
 * Total Accepted:    44.3K
 * Total Submissions: 96.5K
 * Testcase Example:  '"abpcplea"\n["ale","apple","monkey","plea"]'
 *
 * 
 * Given a string and a string dictionary, find the longest string in the
 * dictionary that can be formed by deleting some characters of the given
 * string. If there are more than one possible results, return the longest word
 * with the smallest lexicographical order. If there is no possible result,
 * return the empty string.
 * 
 * Example 1:
 * 
 * Input:
 * s = "abpcplea", d = ["ale","apple","monkey","plea"]
 * 
 * Output: 
 * "apple"
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * s = "abpcplea", d = ["a","b","c"]
 * 
 * Output: 
 * "a"
 * 
 * 
 * 
 * Note:
 * 
 * All the strings in the input will only contain lower-case letters.
 * The size of the dictionary won't exceed 1,000.
 * The length of all the strings in the input won't exceed 1,000.
 * 
 * 
 */
func findLongestWord(s string, d []string) string {
	longestWord := "" // 初始化最长子字符串 
	// 判断 go 的字典序--直接比较
	for _, target := range d {
		l1 := len(longestWord)
		l2 := len(target)
		if l1 > l2 || (l1==l2 && longestWord < target) { // 目标字符串长度不够或者字典序更大
			continue
		}
		if isSubStr(s, target) {
			longestWord = target			
		}
	}
	return longestWord
}

// 采用双指针判断是否是子字符串
func isSubStr(s, target string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(target){
		if s[i] == target[j] {
			j ++
		}
		i ++
	}
	// 如果是字符串，则 target 遍历完成，而 j 多自加了一次，所以应该相等 ss
	return j == len(target)
}

