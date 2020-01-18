/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 *
 * https://leetcode.com/problems/number-of-boomerangs/description/
 *
 * algorithms
 * Easy (50.66%)
 * Likes:    345
 * Dislikes: 554
 * Total Accepted:    62K
 * Total Submissions: 121.7K
 * Testcase Example:  '[[0,0],[1,0],[2,0]]'
 *
 * Given n points in the plane that are all pairwise distinct, a "boomerang" is
 * a tuple of points (i, j, k) such that the distance between i and j equals
 * the distance between i and k (the order of the tuple matters).
 * 
 * Find the number of boomerangs. You may assume that n will be at most 500 and
 * coordinates of points are all in the range [-10000, 10000] (inclusive).
 * 
 * Example:
 * 
 * 
 * Input:
 * [[0,0],[1,0],[2,0]]
 * 
 * Output:
 * 2
 * 
 * Explanation:
 * The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
 * 
 * 
 * 
 * 
 */

// @lc code=start
//计算数组能组成的回旋镖个数
// 31/31 cases passed (204 ms)
// Your runtime beats 33.33 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.5 MB)
func numberOfBoomerangs(points [][]int) int {
	var countMap map[int]int
	var res int
    for _, point := range points {
		countMap = make(map[int]int)
		//统计到当前点距离相同的所有点,根据距离不同放入不同的 map 键中
		for _, point2 := range points {
			countMap[(point[0]-point2[0]) * (point[0]-point2[0]) + (point[1]-point2[1]) * (point[1]-point2[1])] += 1
		}
		for _, count := range countMap {
			if count >= 2 {
				res += count * (count - 1)
			}
		}
	}
	return res
}
// @lc code=end

