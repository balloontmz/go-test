/*
 * @lc app=leetcode id=407 lang=golang
 *
 * [407] Trapping Rain Water II
 *
 * https://leetcode.com/problems/trapping-rain-water-ii/description/
 *
 * algorithms
 * Hard (40.18%)
 * Likes:    953
 * Dislikes: 25
 * Total Accepted:    34.2K
 * Total Submissions: 84.3K
 * Testcase Example:  '[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]'
 *
 * Given an m x n matrix of positive integers representing the height of each
 * unit cell in a 2D elevation map, compute the volume of water it is able to
 * trap after raining.
 * 
 * 
 * 
 * Note:
 * 
 * Both m and n are less than 110. The height of each unit cell is greater than
 * 0 and is less than 20,000.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given the following 3x6 height map:
 * [
 * ⁠ [1,4,3,1,3,2],
 * ⁠ [3,2,1,3,2,4],
 * ⁠ [2,3,3,2,3,1]
 * ]
 * 
 * Return 4.
 * 
 * 
 * 
 * 
 * The above image represents the elevation map
 * [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] before the rain.
 * 
 * 
 * 
 * 
 * 
 * After the rain, water is trapped between the blocks. The total volume of
 * water trapped is 4.
 * 
 */

// @lc code=start
//采用优先队列解决 3d 的最大蓄水量
// 40/40 cases passed (40 ms)
// Your runtime beats 22.22 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.2 MB)
import "container/heap"

var Directions = [][]int{
	[]int{1, 0},
	[]int{-1, 0},
	[]int{0, 1},
	[]int{0, -1},
}

func trapRainWater(heightMap [][]int) int {
	//TODO: 关于 row 和 column 的问题一定要注意!!!
	var pq = make(PriorityQueue, 0) // 优先队列,用于存放需要遍历的点
	heap.Init(&pq)

	if len(heightMap) <= 2 || len(heightMap[0]) <= 2 {
		return 0
	}

	var row = len(heightMap)
	var column = len(heightMap[0])

	//初始化 visited map 并将最外层加入其中
	//同时将最外层节点加入优先队列
    var visited = make(map[int](map[int]bool))

	for i := 0; i < row; i++ {
		visited[i] = make(map[int]bool)
		heap.Push(&pq, &Node{
			0, i, heightMap[i][0],
		})
		visited[i][0] = true

		heap.Push(&pq, &Node{
			column-1, i, heightMap[i][column-1],
		})
		visited[i][column-1] = true
	}

	for j := 1; j < column-1; j++ {
		heap.Push(&pq, &Node{
			j, 0, heightMap[0][j],
		})
		visited[0][j] = true

		heap.Push(&pq, &Node{
			j, row-1, heightMap[row-1][j],
		})
		visited[row-1][j] = true
	}

	var res int
	var mx int
	//遍历所有优先节点
	for pq.Len() != 0 {
		var t = heap.Pop(&pq).(*Node)
        var h = t.height
		var r = t.r
		var c = t.c
		mx = max(mx, h) // 当前海平面高度 -- 采用的优先队列,递增
        
		//遍历每个方向,如果当前遍历的点能进水,则能蓄水
		for _, d := range Directions {
			var x = r + d[0]
			var y = c + d[1]
			if x < 0 || x >= row || y < 0 || y >= column || visited[x][y] {
				continue
			}
            // if _, ok := visited[x][y]; ok {
			// 	continue
			// }
			visited[x][y] = true
			if heightMap[x][y] < mx {
                // fmt.Print("当前添加的数为：", mx - heightMap[x][y], "当前高度为：", heightMap[x][y], "\n")
				res += mx - heightMap[x][y]
			}
			//将当前遍历的点放入队列
			heap.Push(&pq, &Node{
				y, x, heightMap[x][y],
			})
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

//采用标准库的 heap 实现优先队列
type Node struct {
	c, r, height int
}

type PriorityQueue []*Node

//实现 sort 接口
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//升序排列
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq PriorityQueue) Swap(i, j int)  {
	pq[i], pq[j] = pq[j], pq[i]
}

//实现堆接口
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// @lc code=end

