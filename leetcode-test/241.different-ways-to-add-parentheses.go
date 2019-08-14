/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 *
 * https://leetcode.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (50.88%)
 * Likes:    1030
 * Dislikes: 52
 * Total Accepted:    78.6K
 * Total Submissions: 154.2K
 * Testcase Example:  '"2-1-1"'
 *
 * Given a string of numbers and operators, return all possible results from
 * computing all the different possible ways to group numbers and operators.
 * The valid operators are +, - and *.
 * 
 * Example 1:
 * 
 * 
 * Input: "2-1-1"
 * Output: [0, 2]
 * Explanation: 
 * ((2-1)-1) = 0 
 * (2-(1-1)) = 2
 * 
 * Example 2:
 * 
 * 
 * Input: "2*3-4*5"
 * Output: [-34, -14, -10, -10, 10]
 * Explanation: 
 * (2*(3-(4*5))) = -34 
 * ((2*3)-(4*5)) = -14 
 * ((2*(3-4))*5) = -10 
 * (2*((3-4)*5)) = -10 
 * (((2*3)-4)*5) = 10
 * 
 */
func diffWaysToCompute(input string) []int {
	// ✔ 25/25 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (6.5 MB)
	// 分冶法：将问题转换为类型相同的子问题。直到子问题无法再次分解。
	var way []int
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '+' || c == '-' || c== '*' {
			left := diffWaysToCompute(input[0:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch c {
						case '+':
							way = append(way, l+r)
						case '-':
							way = append(way, l-r)
						case '*':
							way = append(way, l*r)
					}
				}
			}
		}
	}
	// 对于最小子问题的解决方法
	if len(way) == 0 {
		v, _ := strconv.Atoi(input)
		way = append(way, v)
	}
	return way
}

