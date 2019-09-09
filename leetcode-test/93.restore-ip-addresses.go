/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.10%)
 * Likes:    757
 * Dislikes: 316
 * Total Accepted:    150.2K
 * Total Submissions: 465.5K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 * 
 * Example:
 * 
 * 
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 * 
 * 
 */
func restoreIpAddresses(s string) []string {
	// ✔ 147/147 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.1 MB)
	addresses := &[]string{}
	tempAddress := ""
	doRestore(0, tempAddress, addresses, s)
	return *addresses
}

// doRestore k 代表遍历深度
func doRestore(k int, tempAddress string, addresses *[]string, s string)  {
	// 叶子节点标志
	if k == 4 || len(s) == 0 {
		if k == 4 && len(s) == 0 {
			*addresses = append(*addresses, tempAddress)
		}
		return
	}

	// i < len(s) 的意义？
	for i := 0; i < len(s) && i <= 2; i++ {
		if i != 0 && s[0] == '0' {
			break
		}
		part := s[:i+1]
		num, _ := strconv.Atoi(part)
		if num <= 255 {
			if len(tempAddress) != 0 {
				part = "." + part
			}
			tempAddress = tempAddress + part
			doRestore(k+1, tempAddress, addresses, s[i+1:])
			// 回溯算法典型标志 -- 递归返回时临时变量还原
			// fmt.Print("删除前的 temp 为：", tempAddress)
			tempAddress = tempAddress[:len(tempAddress) - len(part)] 
			// fmt.Print("删除后的 temp 为：", tempAddress)
		}
	}
}

