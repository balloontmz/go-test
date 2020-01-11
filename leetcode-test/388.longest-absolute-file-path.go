/*
 * @lc app=leetcode id=388 lang=golang
 *
 * [388] Longest Absolute File Path
 *
 * https://leetcode.com/problems/longest-absolute-file-path/description/
 *
 * algorithms
 * Medium (40.13%)
 * Likes:    526
 * Dislikes: 1172
 * Total Accepted:    77K
 * Total Submissions: 189.3K
 * Testcase Example:  '"dir\\n\\tsubdir1\\n\\tsubdir2\\n\\t\\tfile.ext"'
 *
 * Suppose we abstract our file system by a string in the following manner:
 * 
 * The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:
 * 
 * dir
 * ⁠   subdir1
 * ⁠   subdir2
 * ⁠       file.ext
 * 
 * 
 * The directory dir contains an empty sub-directory subdir1 and a
 * sub-directory subdir2 containing a file file.ext.
 * 
 * The string
 * "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
 * represents:
 * 
 * dir
 * ⁠   subdir1
 * ⁠       file1.ext
 * ⁠       subsubdir1
 * ⁠   subdir2
 * ⁠       subsubdir2
 * ⁠           file2.ext
 * 
 * 
 * The directory dir contains two sub-directories subdir1 and subdir2. subdir1
 * contains a file file1.ext and an empty second-level sub-directory
 * subsubdir1. subdir2 contains a second-level sub-directory subsubdir2
 * containing a file file2.ext.
 * 
 * We are interested in finding the longest (number of characters) absolute
 * path to a file within our file system. For example, in the second example
 * above, the longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and
 * its length is 32 (not including the double quotes).
 * 
 * Given a string representing the file system in the above format, return the
 * length of the longest absolute path to file in the abstracted file system.
 * If there is no file in the system, return 0.
 * 
 * Note:
 * 
 * The name of a file contains at least a . and an extension.
 * The name of a directory or sub-directory will not contain a ..
 * 
 * 
 * 
 * Time complexity required: O(n) where n is the size of the input string.
 * 
 * Notice that a/aa/aaa/file1.txt is not the longest file path, if there is
 * another path aaaaaaaaaaaaaaaaaaaaa/sth.png.
 */

// @lc code=start
//寻找最长的文件路径
// 25/25 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
func lengthLongestPath(input string) int {
	var m = make(map[int]int)
	var res = 0
	var level = 0
	var l = len(input)


	for i := 0; i < l; i++ {
		var start = i // 指定开始值
		//遍历到结束位置
		for i < l && input[i] != '\t' && input[i] != '\n' {
			i++
		}
		if i >= l || input[i] == '\n'  {
			var tag = false
			for j := start; j < i; j++ {
				if input[j] == '.' {
					res = max(res, m[level] + i - start)
					tag = true
					break;
				}
			}
			if !tag {
				level++
				m[level] = m[level-1] + i-start+1		
			}
			level = 0
		} else {
			level++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

