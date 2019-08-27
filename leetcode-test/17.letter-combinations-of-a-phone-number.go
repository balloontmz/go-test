/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (42.47%)
 * Likes:    2493
 * Dislikes: 328
 * Total Accepted:    432.9K
 * Total Submissions: 1M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 * 
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * Note:
 * 
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 * 
 */
var KEYS = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	// TODO: 此处出现过由于传递的不是切片的指针而没有改变切片的值的情况
	// ✔ 25/25 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 9.09 % of golang submissions (2.6 MB)
	var combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}
	doCombination("", &combinations, digits)
	return combinations
}

// 传递指针，方便递归修改切片内的值
func doCombination(s string, combinations *[]string, digits string)  {
	if len(s) == len(digits) {
		*combinations = append(*combinations, s)
		// fmt.Println("进入此处", combinations)
		return
	}
	// fmt.Println("当前遍历的索引为：", len(s), len(digits))
	// digits[len(s)] 为 int32  rune 类型
	// '0' 为  uint8  byte 类型
	// 此处为什么能直接相减呢？？？
	var curDigits = digits[len(s)] - '0'
	// fmt.Println("当前遍历的 digits 为：", curDigits)
	var letters = KEYS[curDigits]
	for i := 0; i < len(letters); i++ {
		s += (string)(letters[i])
		// fmt.Println("当前遍历的 letters 为：", letters, "s 为：", s)
		doCombination(s, combinations, digits)
		s = s[:len(s)-1]
	}
}
