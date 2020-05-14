/*
 * @lc app=leetcode id=676 lang=golang
 *
 * [676] Implement Magic Dictionary
 *
 * https://leetcode.com/problems/implement-magic-dictionary/description/
 *
 * algorithms
 * Medium (53.55%)
 * Likes:    522
 * Dislikes: 124
 * Total Accepted:    36.4K
 * Total Submissions: 67.8K
 * Testcase Example:  '["MagicDictionary", "buildDict", "search", "search", "search", "search"]\n' +
  '[[], [["hello","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]'
 *
 * 
 * Implement a magic directory with buildDict, and search methods.
 * 
 * 
 * 
 * For the method buildDict, you'll be given a list of non-repetitive words to
 * build a dictionary.
 * 
 * 
 * 
 * For the method search, you'll be given a word, and judge whether if you
 * modify exactly one character into another character in this word, the
 * modified word is in the dictionary you just built.
 * 
 * 
 * Example 1:
 * 
 * Input: buildDict(["hello", "leetcode"]), Output: Null
 * Input: search("hello"), Output: False
 * Input: search("hhllo"), Output: True
 * Input: search("hell"), Output: False
 * Input: search("leetcoded"), Output: False
 * 
 * 
 * 
 * Note:
 * 
 * You may assume that all the inputs are consist of lowercase letters a-z.
 * For contest purpose, the test data is rather small by now. You could think
 * about highly efficient algorithm after the contest.
 * Please remember to RESET your class variables declared in class
 * MagicDictionary, as static/class variables are persisted across multiple
 * test cases. Please see here for more details.
 * 
 * 
 */

// @lc code=start
// 32/32 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.2 MB)
type MagicDictionary struct {
    data map[int]([]string)
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    var data = make(map[int]([]string), 0)
    return MagicDictionary{
        data: data,
    }
}


/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string)  {
    for _, str := range dict {
        this.data[len(str)] = append(this.data[len(str)], str)
    }
    return
}


/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
    var data = this.data[len(word)]
    if data == nil {
        return false
    }
    for _, str := range data {
        if verifyMagic(str, word) {
            return true
        }
    }
    return false
}

func verifyMagic(x, y string) bool  {
    var cnt = 0
    for i := 0; i < len(x); i++ {
        if x[i] != y[i] {
            cnt++
        }
    }
    return cnt == 1
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
// @lc code=end

