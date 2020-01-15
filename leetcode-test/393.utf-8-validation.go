/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 *
 * https://leetcode.com/problems/utf-8-validation/description/
 *
 * algorithms
 * Medium (36.55%)
 * Likes:    181
 * Dislikes: 859
 * Total Accepted:    43.3K
 * Total Submissions: 117.6K
 * Testcase Example:  '[197,130,1]'
 *
 * A character in UTF8 can be from 1 to 4 bytes long, subjected to the
 * following rules:
 * 
 * For 1-byte character, the first bit is a 0, followed by its unicode code.
 * For n-bytes character, the first n-bits are all one's, the n+1 bit is 0,
 * followed by n-1 bytes with most significant 2 bits being 10.
 * 
 * This is how the UTF-8 encoding would work:
 * 
 * ⁠  Char. number range  |        UTF-8 octet sequence
 * ⁠     (hexadecimal)    |              (binary)
 * ⁠  --------------------+---------------------------------------------
 * ⁠  0000 0000-0000 007F | 0xxxxxxx
 * ⁠  0000 0080-0000 07FF | 110xxxxx 10xxxxxx
 * ⁠  0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
 * ⁠  0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
 * 
 * 
 * Given an array of integers representing the data, return whether it is a
 * valid utf-8 encoding.
 * 
 * 
 * Note:
 * The input is an array of integers. Only the least significant 8 bits of each
 * integer is used to store the data. This means each integer represents only 1
 * byte of data.
 * 
 * 
 * 
 * Example 1:
 * 
 * data = [197, 130, 1], which represents the octet sequence: 11000101 10000010
 * 00000001.
 * 
 * Return true.
 * It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte
 * character.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * data = [235, 140, 4], which represented the octet sequence: 11101011
 * 10001100 00000100.
 * 
 * Return false.
 * The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes
 * character.
 * The next byte is a continuation byte which starts with 10 and that's
 * correct.
 * But the second continuation byte does not start with 10, so it is invalid.
 * 
 * 
 */

// @lc code=start
//验证是否为有效的 utf8 编码
// 49/49 cases passed (12 ms)
// Your runtime beats 73.08 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5.4 MB)
//修剪代码之后的结果 -- 
// 49/49 cases passed (12 ms)
// Your runtime beats 73.08 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5.4 MB)
func validUtf8(data []int) bool {
	var cnt int
	for _, v := range data {
		if cnt == 0 {
			if (v >> 5) == 6 {
				cnt = 1
			} else if (v >> 4) == 14 {
				cnt = 2
			} else if (v >> 3) == 30 {
				cnt = 3
			} else if (v >> 7) == 1 {
				return false
			}
		} else {
			if (v >> 6) != 2 {
				return false
			}
			cnt--
		}
	}
	return cnt == 0 // 此处 cnt不为 0 也是错误的
}
// func validUtf8(data []int) bool {
// 	var n = len(data)
// 	for i := 0; i < n; i++ {
// 		if data[i] < 128 {
// 			continue
// 		} else {
// 			var cnt = 0
// 			var val = data[i]
// 			for j := 7; j >= 1; j-- {
// 				if val >= Pow(2, j) {
// 					cnt++
// 				} else {
// 					break
// 				}
// 				val -= Pow(2, j)
// 			}
// 			if cnt == 1 || cnt > 4 || cnt > n-i  {
// 				return false
// 			}
// 			for j := i + 1; j < i + cnt; j++ {
// 				if (data[j] > 191 || data[j] < 128) {
// 					return false
// 				} 
// 			} 
// 			i += cnt - 1
// 		}
// 	}
// 	return true
// }

// func Pow(x, y int) int {
// 	var res = 1
// 	for i := 0; i < y; i++ {
// 		res *= x
// 	}
// 	return res
// }
// @lc code=end

