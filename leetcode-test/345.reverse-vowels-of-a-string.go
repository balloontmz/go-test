/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (41.53%)
 * Likes:    389
 * Dislikes: 718
 * Total Accepted:    158.1K
 * Total Submissions: 380.2K
 * Testcase Example:  '"hello"'
 *
 * Write a function that takes a string as input and reverse only the vowels of
 * a string.
 * 
 * Example 1:
 * 
 * 
 * Input: "hello"
 * Output: "holle"
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "leetcode"
 * Output: "leotcede"
 * 
 * 
 * Note:
 * The vowels does not include the letter "y".
 * 
 * 
 * 
 */
 // 双指针，翻转元音字符。
 import (
	 "strings"
 )

 var vowels = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}


func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1
	res := make([]byte, len(s))
	for i <= j {
		ci := s[i]
		cj := s[j]
		if !contain(ci, vowels) {
			res[i] = ci
			i++
		} else if !contain(cj, vowels){
			res[j] = cj
			j--
		} else {
			res[i] = cj
			res[j] = ci
			i++
			j--
		}
	}
	return string(res)
}

func contain(s byte, vow []byte) bool {
	for _, val := range vow {
		if val == s {
			return true
		}
	}
	return false
}

