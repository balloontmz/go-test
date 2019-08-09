/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 *
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (44.35%)
 * Likes:    239
 * Dislikes: 371
 * Total Accepted:    46K
 * Total Submissions: 103.7K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 * 
 * Given a list of sorted characters letters containing only lowercase letters,
 * and given a target letter target, find the smallest element in the list that
 * is larger than the given target.
 * 
 * Letters also wrap around.  For example, if the target is target = 'z' and
 * letters = ['a', 'b'], the answer is 'a'.
 * 
 * 
 * Examples:
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 * 
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 * 
 * 
 * 
 * Note:
 * 
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique
 * letters.
 * target is a lowercase letter.
 * 
 * 
 */
func nextGreatestLetter(letters []byte, target byte) byte {
	// ✔ 165/165 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.9 MB)
	// if len(letters) == 0 || (len(letters) > 0 && (letters[len(letters)-1] < target || letters[0] > target)) {
	// 	return letters[0]
	// }
	var l = 0
	var h = len(letters) - 1
	for l <= h {
		m := l + (h-l)/2
        if (letters[m] <= target) {
            l = m + 1;
        } else {
            h = m - 1
        }
	}
	if l >= len(letters) {
		return letters[0]
	}
    return letters[l];
}

