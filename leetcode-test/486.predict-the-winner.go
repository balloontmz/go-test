/*
 * @lc app=leetcode id=486 lang=golang
 *
 * [486] Predict the Winner
 *
 * https://leetcode.com/problems/predict-the-winner/description/
 *
 * algorithms
 * Medium (47.12%)
 * Likes:    1257
 * Dislikes: 79
 * Total Accepted:    62K
 * Total Submissions: 131K
 * Testcase Example:  '[1,5,2]'
 *
 * Given an array of scores that are non-negative integers. Player 1 picks one
 * of the numbers from either end of the array followed by the player 2 and
 * then player 1 and so on. Each time a player picks a number, that number will
 * not be available for the next player. This continues until all the scores
 * have been chosen. The player with the maximum score wins. 
 * 
 * Given an array of scores, predict whether player 1 is the winner. You can
 * assume each player plays to maximize his score. 
 * 
 * Example 1:
 * 
 * Input: [1, 5, 2]
 * Output: False
 * Explanation: Initially, player 1 can choose between 1 and 2. If he chooses 2
 * (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5,
 * then player 1 will be left with 1 (or 2). So, final score of player 1 is 1 +
 * 2 = 3, and player 2 is 5. Hence, player 1 will never be the winner and you
 * need to return False.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [1, 5, 233, 7]
 * Output: True
 * Explanation: Player 1 first chooses 1. Then player 2 have to choose between
 * 5 and 7. No matter which number player 2 choose, player 1 can choose
 * 233.Finally, player 1 has more score (234) than player 2 (12), so you need
 * to return True representing player1 can win.
 * 
 * 
 * 
 * Note:
 * 
 * 1 
 * Any scores in the given array are non-negative integers and will not exceed
 * 10,000,000.
 * If the scores of both players are equal, then player 1 is still the winner.
 * 
 * 
 */

// @lc code=start
//递归解法,还可以辅助采用动态规划
// 61/61 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
func PredictTheWinner(nums []int) bool {
	//假设每个人都想赢
	return judgeWinner(1, nums, 0, 0)
}

func judgeWinner(player int, nums []int, nums1, nums2 int) bool {
	// fmt.Print("当前传入的 nums1 为:", nums1, "nums2 为:", nums2)
	if len(nums) == 0 {
		return nums1 >= nums2
	}
	if len(nums) == 1 {
		if player == 1 {
			return nums1 + nums[0] >= nums2
		} else if player == 2 {
			//等于的情况,第一位玩家获胜.
			return nums2 + nums[0] > nums1
		}
	}
	numsF := nums[1:]
	numsB := nums[:len(nums)-1]
	if player == 1 {
		return !judgeWinner(2, numsF, nums1+nums[0], nums2) || !judgeWinner(2, numsB, nums1+nums[len(nums)-1], nums2)
	}
	return !judgeWinner(1, numsF, nums1, nums2+nums[0]) || !judgeWinner(1, numsB, nums1, nums2+nums[len(nums)-1])
}
// @lc code=end

