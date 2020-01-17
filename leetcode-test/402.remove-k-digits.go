/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 *
 * https://leetcode.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (27.13%)
 * Likes:    1294
 * Dislikes: 76
 * Total Accepted:    81.4K
 * Total Submissions: 297.4K
 * Testcase Example:  '"1432219"\n3'
 *
 * Given a non-negative integer num represented as a string, remove k digits
 * from the number so that the new number is the smallest possible.
 * 
 * 
 * Note:
 * 
 * The length of num is less than 10002 and will be ≥ k.
 * The given num does not contain any leading zero.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * Input: num = "1432219", k = 3
 * Output: "1219"
 * Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219
 * which is the smallest.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: num = "10200", k = 1
 * Output: "200"
 * Explanation: Remove the leading 1 and the number is 200. Note that the
 * output must not contain leading zeroes.
 * 
 * 
 * 
 * Example 3:
 * 
 * Input: num = "10", k = 2
 * Output: "0"
 * Explanation: Remove all the digits from the number and it is left with
 * nothing which is 0.
 * 
 * 
 */

// @lc code=start
//删除 k 个字符
//维护一个单调递增栈,只到 k 等于 0 或者遍历完成,如果是遍历完成,则遗弃后方的字节
// 33/33 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.7 MB)
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack []byte
	for i := 0; i < len(num); i++ {
		// if i + k == len(num) {
		// 	return string(stack)
		// } else
		for len(stack) != 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
    
    if k != 0 {
        return string(stack[:len(stack)-k])
    }

	for k, v := range stack {
		if v != '0' {
			return string(stack[k:])
		}
	}
	return "0"
}
// @lc code=end

