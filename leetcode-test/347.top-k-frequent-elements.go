/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (55.05%)
 * Likes:    1682
 * Dislikes: 101
 * Total Accepted:    224.2K
 * Total Submissions: 401.8K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given a non-empty array of integers, return the k most frequent elements.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1], k = 1
 * Output: [1]
 * 
 * 
 * Note: 
 * 
 * 
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Your algorithm's time complexity must be better than O(n log n), where n is
 * the array's size.
 * 
 * 
 */
 // 桶排序 取频率最多的 k 个元素
func topKFrequent(nums []int, k int) []int {
	frequencyForNum := make(map[int]int)
	// fmt.Print("不存在的键取出的值为：", frequencyForNum[0]) -- 0
	for _, v := range nums {
		frequencyForNum[v] = frequencyForNum[v] + 1
	}
	buckets := make(map[int]([]int))
	for k, v := range frequencyForNum { // go 的 map 是无序的！！！
		// fmt.Print("测试初始化的 buckets[v] 值为：", buckets[v])
		buckets[v] = append(buckets[v], k)  // 此处应该会报错，报错原因为 slice 未初始化
	}
	topK := make([]int, 0)
	// fmt.Print("组合的 map 为：", buckets)
	for i := len(nums); i >= 0; i-- {
		if len(topK) >= k {
			break
		}
		if buckets[i] == nil { // 如果键对应的值不存在，进行下次遍历
			continue
		}
		// topK = append(topK, buckets[i]...) // 此处会超出数组长度
		if len(buckets[i]) <= (k - len(topK)) {
			topK = append(topK, buckets[i]...)
		} else {
			topK = append(topK, buckets[i][:(k - len(topK))]...)
		}
	}

	return topK
}

