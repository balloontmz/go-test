/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 *
 * https://leetcode.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (45.90%)
 * Likes:    459
 * Dislikes: 797
 * Total Accepted:    74K
 * Total Submissions: 160.2K
 * Testcase Example:  '0'
 *
 * A binary watch has 4 LEDs on the top which represent the hours (0-11), and
 * the 6 LEDs on the bottom represent the minutes (0-59).
 * Each LED represents a zero or one, with the least significant bit on the
 * right.
 * 
 * For example, the above binary watch reads "3:25".
 * 
 * Given a non-negative integer n which represents the number of LEDs that are
 * currently on, return all possible times the watch could represent.
 * 
 * Example:
 * Input: n = 1Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04",
 * "0:08", "0:16", "0:32"]
 * 
 * 
 * Note:
 * 
 * The order of output does not matter.
 * The hour must not contain a leading zero, for example "01:00" is not valid,
 * it should be "1:00".
 * The minute must be consist of two digits and may contain a leading zero, for
 * example "10:2" is not valid, it should be "10:02".
 * 
 * 
 */

// @lc code=start
//二进制表
// 10/10 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.9 MB)
var hour = []int{
	8, 4, 2, 1,
}

var minute = []int{
	32,16,8,4,2,1,
}

func readBinaryWatch(num int) []string {
	var res []string
	for i := 0; i <= num; i++ {
		var hours = generate(hour, i)
		var minutes = generate(minute, num-i)
		for _, h := range hours {
            if h >= 12 {
                continue
            }
			var	pre = strconv.Itoa(h)
			for _, m := range minutes {
				if m > 59 {
					continue
				}
				if m < 10 {
					res = append(res, pre + ":0" + strconv.Itoa(m))
				} else {
					res = append(res, pre + ":" + strconv.Itoa(m))
				}
			}
		}
	}
	return res
}

func generate(seed []int, cnt int) []int {
	var res = make([]int, 0)
	helper(&res, cnt, 0, 0, seed)
	return res
}

func helper(out *[]int, cnt, pos, sum int, set []int)  {
	if cnt == 0 {
		*out = append(*out, sum)
	}
	for i := pos; i < len(set); i++ {
		helper(out, cnt-1, i+1, sum + set[i], set)
	}
}
// @lc code=end

