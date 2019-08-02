/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 *
 * https://leetcode.com/problems/can-place-flowers/description/
 *
 * algorithms
 * Easy (30.93%)
 * Likes:    498
 * Dislikes: 277
 * Total Accepted:    66.4K
 * Total Submissions: 213.3K
 * Testcase Example:  '[1,0,0,0,1]\n1'
 *
 * Suppose you have a long flowerbed in which some of the plots are planted and
 * some are not. However, flowers cannot be planted in adjacent plots - they
 * would compete for water and both would die.
 * 
 * Given a flowerbed (represented as an array containing 0 and 1, where 0 means
 * empty and 1 means not empty), and a number n, return if n new flowers can be
 * planted in it without violating the no-adjacent-flowers rule.
 * 
 * Example 1:
 * 
 * Input: flowerbed = [1,0,0,0,1], n = 1
 * Output: True
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: flowerbed = [1,0,0,0,1], n = 2
 * Output: False
 * 
 * 
 * 
 * Note:
 * 
 * The input array won't violate no-adjacent-flowers rule.
 * The input array size is in the range of [1, 20000].
 * n is a non-negative integer which won't exceed the input array size.
 * 
 * 
 */
func canPlaceFlowers(flowerbed []int, n int) bool {
	// ✔ 123/123 cases passed (16 ms)
	// ✔ Your runtime beats 92.63 % of golang submissions
	// ✔ Your memory usage beats 50 % of golang submissions (5.9 MB)
	// 计算花床还能种下多少朵花
	var l int  = len(flowerbed)
	var cnt = 0
	for i := 0; i < l; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		var pre, next int  // 初始化
		if i != 0 {
			pre = flowerbed[i-1]
		} 
		if i != l - 1 {
			next = flowerbed[i+1]
		}
		if pre == 0 && next == 0 {
			// 已经种下的花不算
			cnt++
			flowerbed[i] =1;
		}
	}
	return cnt >= n
}

