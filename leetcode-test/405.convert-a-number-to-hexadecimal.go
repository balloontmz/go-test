/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 *
 * https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
 *
 * algorithms
 * Easy (42.68%)
 * Likes:    323
 * Dislikes: 86
 * Total Accepted:    54.5K
 * Total Submissions: 127.6K
 * Testcase Example:  '26'
 *
 *
 * Given an integer, write an algorithm to convert it to hexadecimal. For
 * negative integer, two’s complement method is used.
 *
 *
 * Note:
 *
 * All letters in hexadecimal (a-f) must be in lowercase.
 * The hexadecimal string must not contain extra leading 0s. If the number is
 * zero, it is represented by a single zero character '0'; otherwise, the first
 * character in the hexadecimal string will not be the zero character.
 * The given number is guaranteed to fit within the range of a 32-bit signed
 * integer.
 * You must not use any method provided by the library which converts/formats
 * the number to hex directly.
 *
 *
 *
 * Example 1:
 *
 * Input:
 * 26
 *
 * Output:
 * "1a"
 *
 *
 *
 * Example 2:
 *
 * Input:
 * -1
 *
 * Output:
 * "ffffffff"
 *
 *
 */

// @lc code=start
//16进制转换利用go数据类型特性的算法
func toHex(num int) string {
	// 100/100 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	if num == 0 {
		return "0"
	}
	var res = ""
	var temp = "0123456789abcdef"
	for num != 0 {
		res = string(temp[num&15]) + res
		num = int(uint32(num) / 16)
	}
	return res
}

//16进制转换正常算法
// func toHex(num int) string {
// 	// 100/100 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (1.9 MB)
// 	var res = ""
// 	var temp = "0123456789abcdef"
// 	if num == 0 {
// 		return "0"
// 	}
// 	if num < 0 {
// 		num = -num
// 		// 计算补码
// 		num ^= 0xffffffff
// 		num += 1
// 	}
// 	for num > 0 {
// 		res = string(temp[num%16]) + res
// 		num = num / 16
// 	}
// 	return res
// }

// @lc code=end

