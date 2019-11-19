import "fmt"

/*
 * @lc app=leetcode id=677 lang=golang
 *
 * [677] Map Sum Pairs
 *
 * https://leetcode.com/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (52.42%)
 * Likes:    363
 * Dislikes: 64
 * Total Accepted:    31.4K
 * Total Submissions: 59.9K
 * Testcase Example:  '["MapSum", "insert", "sum", "insert", "sum"]\n' +
  '[[], ["apple",3], ["ap"], ["app",2], ["ap"]]'
 *
 *
 * Implement a MapSum class with insert, and sum methods.
 *
 *
 *
 * For the method insert, you'll be given a pair of (string, integer). The
 * string represents the key and the integer represents the value. If the key
 * already existed, then the original key-value pair will be overridden to the
 * new one.
 *
 *
 *
 * For the method sum, you'll be given a string representing the prefix, and
 * you need to return the sum of all the pairs' value whose key starts with the
 * prefix.
 *
 *
 * Example 1:
 *
 * Input: insert("apple", 3), Output: Null
 * Input: sum("ap"), Output: 3
 * Input: insert("app", 2), Output: Null
 * Input: sum("ap"), Output: 5
 *
 *
 *
*/

// @lc code=start
//实现一个trie用来求前缀和
// 30/30 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.6 MB)
type MapSum struct {
	Childs map[int]*MapSum
	Val    int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func indexForChar(index byte) int {
	return int(index - 'a')
}

func (this *MapSum) Insert(key string, val int) {
	if this == nil {
		return
	}
	if len(key) == 0 {
		this.Val = val
		return
	}
	var index = indexForChar(key[0])
	if this.Childs == nil { // 预期可能需要该代码！ -- 确实需要
		this.Childs = make(map[int]*MapSum, 0)
	}
	if this.Childs[index] == nil {
		this.Childs[index] = &MapSum{}
	}
	this.Childs[index].Insert(key[1:], val)
	return
}

func (this *MapSum) Sum(prefix string) int {
	if this == nil { // 此处是否必须？ -- 必须！！！
		return 0
	}
	if len(prefix) != 0 {
		var index = indexForChar(prefix[0])
		return this.Childs[index].Sum(prefix[1:])
	}
	sum := this.Val
	for _, child := range this.Childs {
		fmt.Println("当前遍历的child为：", child)
		sum += child.Sum(prefix)
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

