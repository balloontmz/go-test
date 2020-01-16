/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 *
 * https://leetcode.com/problems/evaluate-division/description/
 *
 * algorithms
 * Medium (49.17%)
 * Likes:    1765
 * Dislikes: 136
 * Total Accepted:    105.7K
 * Total Submissions: 212.7K
 * Testcase Example:  '[["a","b"],["b","c"]]\n[2.0,3.0]\n[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'
 *
 * Equations are given in the format A / B = k, where A and B are variables
 * represented as strings, and k is a real number (floating point number).
 * Given some queries, return the answers. If the answer does not exist, return
 * -1.0.
 * 
 * Example:
 * Given  a / b = 2.0, b / c = 3.0.
 * queries are:  a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
 * return  [6.0, 0.5, -1.0, 1.0, -1.0 ].
 * 
 * The input is:  vector<pair<string, string>> equations, vector<double>&
 * values, vector<pair<string, string>> queries , where equations.size() ==
 * values.size(), and the values are positive. This represents the equations.
 * Return  vector<double>.
 * 
 * According to the example above:
 * 
 * 
 * equations = [ ["a", "b"], ["b", "c"] ],
 * values = [2.0, 3.0],
 * queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]
 * ]. 
 * 
 * 
 * 
 * The input is always valid. You may assume that evaluating the queries will
 * result in no division by zero and there is no contradiction.
 * 
 */

// @lc code=start
//递归,深度优先遍历,此处也可采用广度优先遍历的迭代法,采用队列存储数据
// 11/11 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
//迭代,采用广度优先遍历 bfs
// 11/11 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.2 MB)
type void struct{}

var voidVar void

var m map[string](map[string]float64)

type pair struct {
	end string
	result float64
}

//迭代 bfs
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m = make(map[string](map[string]float64)) // 先重置存储容器
	//容器中值存储
	for k, equation := range equations {
		if m[equation[0]] == nil {
			m[equation[0]] = make(map[string]float64)
		}
		if m[equation[1]] == nil {
			m[equation[1]] = make(map[string]float64)
		}
		m[equation[0]][equation[1]] = values[k]
		m[equation[1]][equation[0]] = 1 / values[k]
	}	
	var res []float64

	for _, query := range queries {
		if m[query[0]] == nil || m[query[1]] == nil {
			res = append(res, -1.0)
			//当前遍历完成
			continue
		} else if query[0] == query[1] {
			//第一个筛选条件已经过滤了 query 不存在的情况
			res = append(res, 1.0)
			continue
		}
		var visited = make(map[string]void)
		var found = false
		var queue []pair // 注意每次遍历队列初始化
		queue = append(queue, pair{query[0], 1.0})
		for len(queue) != 0 && !found {
			var t = queue[0]
			queue = queue[1:]
			if t.end == query[1] {
				found = true
				res = append(res, t.result)
				break
			}
			for k, v := range m[t.end] {
				if _, ok := visited[k]; ok {
					continue
				}
				visited[k] = voidVar
				queue = append(queue, pair{k, t.result * v}) // 叠加计算 并放入数组
			}
		}
		if !found {
			res = append(res, -1.0)
		}
	}
	return res
}

//递归,dfs
// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
// 	m = make(map[string](map[string]float64))
// 	for k, equation := range equations {
// 		if m[equation[0]] == nil {
// 			m[equation[0]] = make(map[string]float64)
// 		}
// 		if m[equation[1]] == nil {
// 			m[equation[1]] = make(map[string]float64)
// 		}
// 		m[equation[0]][equation[1]] = values[k]
// 		m[equation[1]][equation[0]] = 1 / values[k]
// 	}
// 	var res []float64
// 	for _, query := range queries {
// 		var visited = make(map[string]void) // 每次查询初始化一次 visited 数组,防止参数污染
// 		res = append(res, helper(query[0], query[1], &visited))
// 	}
// 	return res
// }

// func helper(first string, second string, visited *map[string]void) float64 {
// 	if m[first][second] != 0.0 {
// 		return m[first][second]
// 	}
// 	for k, v := range m[first] {
// 		if _, ok := (*visited)[k]; ok {
// 			continue
// 		}
// 		(*visited)[k] = voidVar
// 		var t = helper(k, second, visited)
// 		if t > 0 {
// 			return t * v
// 		}
// 	}
// 	return -1.0
// }
// @lc code=end

