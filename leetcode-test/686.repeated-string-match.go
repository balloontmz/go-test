/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 *
 * https://leetcode.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Easy (32.15%)
 * Likes:    726
 * Dislikes: 696
 * Total Accepted:    88.2K
 * Total Submissions: 274.2K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * Given two strings A and B, find the minimum number of times A has to be
 * repeated such that B is a substring of it. If no such solution, return -1.
 * 
 * For example, with A = "abcd" and B = "cdabcdab".
 * 
 * Return 3, because by repeating A three times (“abcdabcdabcd”), B is a
 * substring of it; and B is not a substring of A repeated two times
 * ("abcdabcd").
 * 
 * Note:
 * The length of A and B will be between 1 and 10000.
 * 
 */

// @lc code=start
//重复字符串匹配
// 55/55 cases passed (16 ms)
// Your runtime beats 45.45 % of golang submissions
// Your memory usage beats 50 % of golang submissions (8.5 MB)
func repeatedStringMatch(A string, B string) int {
	var la = len(A)
	var lb = len(B)
	var origin = A

	var div = lb / la
	if lb % la == 0  {
		div = lb / la - 1
	}
	for i := 0; i < div; i++ {
		A = A + origin
	}
	if contain(A, B) {
		return div + 1
	}
	A = A + origin
	if contain(A, B) {
		return div + 2
	}
	return -1
}

func contain(A, B string) bool {
	for i := 0; i <= len(A) - len(B); i++ {
		if A[i:len(B)+i] == B {
			return true
		}
	}
	return false
}
// @lc code=end

