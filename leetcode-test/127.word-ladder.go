/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (24.94%)
 * Likes:    1692
 * Dislikes: 873
 * Total Accepted:    287K
 * Total Submissions: 1.1M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 * 
 * 
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 * 
 * 
 * Note:
 * 
 * 
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * 
 * Output: 5
 * 
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 * 
 * Output: 0
 * 
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 * 
 * 
 * 
 * 
 * 
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// ✔ 40/40 cases passed (296 ms)
	// ✔ Your runtime beats 31.94 % of golang submissions
	// ✔ Your memory usage beats 50 % of golang submissions (6 MB)
	// 将初始值加入列表中
	wordList = append(wordList, beginWord)  // 根节点放在末尾
	var N = len(wordList)
	var start = N - 1
	var end = 0
	// 寻找是否有目标字符串，没有直接返回 0
	for end < N && (wordList[end] != endWord) {
		end++
	}
	if end == N {
		return 0
	}
	// 生成每个元素相关联的元素列表 -- 保存索引 -- 能提升效率？
	graphic := buildGraphic(wordList)
	return getShortestPath(graphic, start, end)
}

func buildGraphic(wordList []string) map[int]([]int) {
	var N = len(wordList)
	var graphic = make(map[int]([]int), 0)
	for i := 0; i < N; i++ {
		// 数组有多少项 map 就有多少项
		graphic[i] = []int{}
		for j := 0; j < N; j++ {
			if isConnect(wordList[i], wordList[j]) {
				graphic[i] = append(graphic[i], j)
			}
		}
	}
	return graphic
}

// 试一下返回整数能不能被作为判断条件 -- 不能 必须布尔值
func isConnect(s1, s2 string) bool {
	var diffCnt = 0
	// 当 diffCnt 大于 1 时不用再计算
	for i := 0; i < len(s1) && diffCnt <= 1; i++ {
		if s1[i] != s2[i] {
			diffCnt ++;
		}
	}
	return diffCnt == 1
}

func getShortestPath(graphic map[int]([]int), start, end int) int {
	var queue []int
	var marked = make([]bool, len(graphic))
	queue = append(queue, start)
	// 根节点标记为已遍历
	marked[start] = true
	var path = 1
	// bfs 广度优先遍历核心
	for len(queue) > 0 {
		var size = len(queue)
		path++
		for ;size > 0;size-- {
			var cur = queue[0]
			queue = queue[1:]
			for _, next := range graphic[cur] {
				if next == end {
					return path
				} 
				if marked[next] {  // 如果标记已存在，则跳过加入队列，因为已存在的路径数肯定小于等于当前遍历
					continue
				}
				marked[next] = true
				queue = append(queue, next)
			}
		}
	}
	return 0
}