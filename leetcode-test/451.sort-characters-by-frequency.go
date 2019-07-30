/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 *
 * https://leetcode.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (56.17%)
 * Likes:    772
 * Dislikes: 69
 * Total Accepted:    101K
 * Total Submissions: 178.1K
 * Testcase Example:  '"tree"'
 *
 * Given a string, sort it in decreasing order based on the frequency of
 * characters.
 * 
 * Example 1:
 * 
 * Input:
 * "tree"
 * 
 * Output:
 * "eert"
 * 
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid
 * answer.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * "cccaaa"
 * 
 * Output:
 * "cccaaa"
 * 
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 * 
 * 
 * 
 * Example 3:
 * 
 * Input:
 * "Aabb"
 * 
 * Output:
 * "bbAa"
 * 
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 * 
 * 
 */
func frequencySort(s string) string {
	// ✔ Your runtime beats 46.32 % of golang submissions
	// ✔ Your memory usage beats 54.55 % of golang submissions (5.8 MB)
    // frequencyForNum := make(map[byte]int)
	// var data []byte = []byte(s)
	// // 构建桶
	// for _, v := range data {
	// 	frequencyForNum[v] = frequencyForNum[v] + 1
	// }
	// // 反转桶
	// buckets := make(map[int]([]byte))
	// for k, v := range frequencyForNum { // go 的 map 是无序的！！！
	// 	// fmt.Print("测试初始化的 buckets[v] 值为：", buckets[v])
	// 	buckets[v] = append(buckets[v], k)  // 此处应该会报错，报错原因为 slice 未初始化
	// }
	// // 
	// var res string
	// for i := len(data); i >= 0; i-- {
	// 	if buckets[i] == nil { // 如果键对应的值不存在，进行下次遍历
	// 		continue
	// 	}
	// 	for _, b := range buckets[i] {
	// 		res = res + strings.Repeat(string(b), i)	
	// 	}
	// }

	// return res

	// ✔ 35/35 cases passed (8 ms)
	// ✔ Your runtime beats 80 % of golang submissions
	// ✔ Your memory usage beats 45.45 % of golang submissions (5.8 MB)
	frequencyForNum := make(map[rune]int)
	data := s
	// 构建桶
	for _, v := range data {
		frequencyForNum[v] = frequencyForNum[v] + 1
	}
	// 反转桶
	buckets := make(map[int]([]rune))
	for k, v := range frequencyForNum { // go 的 map 是无序的！！！
		// fmt.Print("测试初始化的 buckets[v] 值为：", buckets[v])
		buckets[v] = append(buckets[v], k)  // 此处应该会报错，报错原因为 slice 未初始化
	}
	// 
	var res string
	for i := len(data); i >= 0; i-- {
		if buckets[i] == nil { // 如果键对应的值不存在，进行下次遍历
			continue
		}
		for _, b := range buckets[i] {
			res = res + strings.Repeat(string(b), i)	
		}
	}

	return res
}

