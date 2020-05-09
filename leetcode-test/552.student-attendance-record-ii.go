/*
 * @lc app=leetcode id=552 lang=golang
 *
 * [552] Student Attendance Record II
 *
 * https://leetcode.com/problems/student-attendance-record-ii/description/
 *
 * algorithms
 * Hard (35.82%)
 * Likes:    497
 * Dislikes: 93
 * Total Accepted:    21.3K
 * Total Submissions: 59.4K
 * Testcase Example:  '1'
 *
 * Given a positive integer n, return the number of all possible attendance
 * records with length n, which will be regarded as rewardable. The answer may
 * be very large, return it after mod 10^9 + 7.
 * 
 * A student attendance record is a string that only contains the following
 * three characters:
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
 * A record is regarded as rewardable if it doesn't contain more than one 'A'
 * (absent) or more than two continuous 'L' (late).
 * 
 * Example 1:
 * 
 * Input: n = 2
 * Output: 8 
 * Explanation:
 * There are 8 records with length 2 will be regarded as rewardable:
 * "PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
 * Only "AA" won't be regarded as rewardable owing to more than one absent
 * times. 
 * 
 * 
 * 
 * Note:
 * The value of n won't exceed 100,000.
 * 
 * 
 * 
 * 
 */

// @lc code=start


func checkRecord(n int) int {
	return checkRecord2(n)
	// return checkRecord1(n)
}

//采用 3 个一维数组存储 dp 结果
// 58/58 cases passed (16 ms)
// Your runtime beats 70 % of golang submissions
// Your memory usage beats 100 % of golang submissions (8.2 MB)
func checkRecord2(n int) int {
	var M = 1000000000 + 7
	var p = make([]int, n)
	var l = make([]int, n)
	var a = make([]int, n)

	p[0] = 1
	l[0] = 1
	a[0] = 1

	//边界条件
	if n > 1 {
		l[1] = 3
		a[1] = 2
	}
	if n > 2 {
		a[2] = 4
	}
	for i := 1; i < n; i++ {
		p[i] = (p[i-1] + l[i-1] + a[i-1]) % M //任何字符最后一位都能接 p
		if i > 1 {
			l[i] = (p[i-1] + a[i-1] + p[i-2] + a[i-2]) % M
		}
		if i > 2 {
			//递推公式
			// a[i] = p1[i-1] + l1[i-1]
			// p1[i] = p1[i-1] + l1[i-1]
			// l1[i] = p1[i-1] + p1[i-2]
			a[i] = (a[i-1] + a[i-2] + a[i-3]) % M
		}
	}
	return (p[n-1] + l[n-1] + a[n-1]) % M
}

//采用一个三维 dp 数组存储结果
// 58/58 cases passed (48 ms)
// Your runtime beats 20 % of golang submissions
// Your memory usage beats 100 % of golang submissions (12.1 MB)
func checkRecord1(n int) int {
	var dp = make([][2][3]int, n+1)
	var M = 1000000000 + 7
	//初始化存储数组 -- 切片性能影响...
	// 58/58 cases passed (524 ms)
	// Your runtime beats 10 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (30.3 MB)
	// for i := 0; i <= n; i++ {
	// 	dpSub := make([][]int, 2)
	// 	dp[i] = dpSub
	// 	for j := 0; j < 2; j++ {
	// 		dpSub[j] = make([]int, 3) // a 不超过 2 连续 k 不超过 3
	// 	}
	// }

	//初始化初始值
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		dp[0][i][j] = 1
	// 	}
	// }
	dp[0] = [2][3]int{
		[3]int{1,1,1},
		[3]int{1,1,1},
	}

	//最多的意思是可以少于
	for i := 1; i <= n; i++ {
		for j := 0; j < 2; j++ { // 最多包含的 j 个数  0 或者 1
			for k := 0; k < 3; k++ { // 最多的包含的连续 L 个数 0 1 2
				val := dp[i-1][j][2] //尾部加 p 的情况
				if j > 0 {
					val += dp[i-1][j-1][2] % M  // 尾部加 a 的情况
				}
				if k > 0 {
					val += dp[i-1][j][k-1] % M// 尾部加 k 的情况
				}
				dp[i][j][k] = val % M
			}
		}
	}
	return dp[n][1][2]
}
// @lc code=end

