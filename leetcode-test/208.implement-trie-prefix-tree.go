/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (41.69%)
 * Likes:    2031
 * Dislikes: 36
 * Total Accepted:    216.7K
 * Total Submissions: 517.6K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Implement a trie with insert, search, and startsWith methods.
 *
 * Example:
 *
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");
 * trie.search("app");     // returns true
 *
 *
 * Note:
 *
 *
 * You may assume that all inputs are consist of lowercase letters a-z.
 * All inputs are guaranteed to be non-empty strings.
 *
 *
*/

// @lc code=start
//实现一个trie 前缀树
// 15/15 cases passed (96 ms)
// Your runtime beats 14.16 % of golang submissions
// Your memory usage beats 100 % of golang submissions (20.2 MB)
type Trie struct {
	Childs map[int]*Trie
	IsLeaf bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if this == nil {
		return
	}
	if len(word) == 0 {
		//代表当前插入已经遍历到了叶子节点
		this.IsLeaf = true
		return
	}
	var index = indexForChar(word[0])
	if this.Childs == nil { // 预期可能需要该代码！ -- 确实需要
		this.Childs = make(map[int]*Trie, 0)
	}
	if this.Childs[index] == nil {
		this.Childs[index] = &Trie{}
	}
	this.Childs[index].Insert(word[1:])
	return
}

func indexForChar(index byte) int {
	return int(index - 'a')
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this == nil {
		return false
	}
	if len(word) == 0 {
		return this.IsLeaf
	}
	var index = indexForChar(word[0])
	return this.Childs[index].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil {
		return false
	}
	if len(prefix) == 0 {
		return true
	}
	var index = indexForChar(prefix[0])
	return this.Childs[index].StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

