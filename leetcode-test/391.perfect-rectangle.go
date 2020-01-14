/*
 * @lc app=leetcode id=391 lang=golang
 *
 * [391] Perfect Rectangle
 *
 * https://leetcode.com/problems/perfect-rectangle/description/
 *
 * algorithms
 * Hard (29.00%)
 * Likes:    288
 * Dislikes: 58
 * Total Accepted:    21.6K
 * Total Submissions: 73.7K
 * Testcase Example:  '[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]'
 *
 * Given N axis-aligned rectangles where N > 0, determine if they all together
 * form an exact cover of a rectangular region.
 * 
 * Each rectangle is represented as a bottom-left point and a top-right point.
 * For example, a unit square is represented as [1,1,2,2]. (coordinate of
 * bottom-left point is (1, 1) and top-right point is (2, 2)).
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4],
 * ⁠ [1,3,2,4],
 * ⁠ [2,3,3,4]
 * ]
 * 
 * Return true. All 5 rectangles together form an exact cover of a rectangular
 * region.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,2,3],
 * ⁠ [1,3,2,4],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4]
 * ]
 * 
 * Return false. Because there is a gap between the two rectangular
 * regions.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [3,2,4,4]
 * ]
 * 
 * Return false. Because there is a gap in the top center.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [2,2,4,4]
 * ]
 * 
 * Return false. Because two of the rectangles overlap with each other.
 * 
 * 
 * 
 */

// @lc code=start
//判断是否为完美矩形
//优化--利用结构体中类型如果可比较则结构体可比较则可用于 map 的键
// 46/46 cases passed (52 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.4 MB)
// 46/46 cases passed (80 ms)
// Your runtime beats 25 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.5 MB)
type void struct{}

const INT_MAX = int(^uint(0) >> 1)

var t void

type point struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {
	

	set := make(map[point]void) // New empty set

	var minX, minY = INT_MAX, INT_MAX
	var MaxX, MaxY int
	var area int

	for _, v := range rectangles {
		minX = min(minX , v[0])
		minY = min(minY , v[1])
		MaxX = max(MaxX , v[2])
		MaxY = max(MaxY , v[3])

		area += (v[3] - v[1]) * (v[2] - v[0])

		for _, x := range []int{v[0], v[2]} {
            for _, y := range []int{v[1], v[3]} {
				p := point{x, y}

				if contain(p, set) {
					delete(set, p)
				} else {
					set[p] = t
				}
			}
		}
	}

	var t1 = point{minX, minY}
	var t2 = point{minX, MaxY}
	var t3 = point{MaxX, minY}
	var t4 = point{MaxX, MaxY}

	if len(set) != 4 {
		return false
	}

	if !contain(t1, set) || !contain(t2, set) || !contain(t3, set) || !contain(t4, set) {
		return false
	}

	return area == (MaxY - minY) * (MaxX - minX)	
}

func contain(s point, m map[point]void) bool {
	_, ok := (m)[s]
	return ok
}

func max(x, y int) int {
	if x>y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x>y {
		return y
	}
	return x
}
// @lc code=end

