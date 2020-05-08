/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 *
 * https://leetcode.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (46.70%)
 * Likes:    246
 * Dislikes: 780
 * Total Accepted:    75.9K
 * Total Submissions: 162.5K
 * Testcase Example:  '"PPALLP"'
 *
 * You are given a string representing an attendance record for a student. The
 * record only contains the following three characters:
 * 
 * 
 * 
 * 'A' : Absent. 
 * 'L' : Late.
 * ⁠'P' : Present. 
 * 
 * 
 * 
 * 
 * A student could be rewarded if his attendance record doesn't contain more
 * than one 'A' (absent) or more than two continuous 'L' (late).    
 * 
 * You need to return whether the student could be rewarded according to his
 * attendance record.
 * 
 * Example 1:
 * 
 * Input: "PPALLP"
 * Output: True
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: "PPALLL"
 * Output: False
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
// 113/113 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
func checkRecord(s string) bool {
	var A int
	var L int
    for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			if A == 1 {
				return false
			}
			A = 1
			L = 0
		} else if s[i] == 'L' {
			if L == 2 {
				return false
			}
			L++
		} else {
			L = 0
		}
	}
	return true
}
// @lc code=end

