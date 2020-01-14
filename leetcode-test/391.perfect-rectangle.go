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
// 46/46 cases passed (80 ms)
// Your runtime beats 25 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.5 MB)
type void struct{}

const INT_MAX = int(^uint(0) >> 1)

var t void

func isRectangleCover(rectangles [][]int) bool {
	

	set := make(map[string]void) // New empty set

	var minX, minY = INT_MAX, INT_MAX
	var MaxX, MaxY int
	var area int

	for _, v := range rectangles {
		minX = min(minX , v[0])
		minY = min(minY , v[1])
		MaxX = max(MaxX , v[2])
		MaxY = max(MaxY , v[3])

		area += (v[3] - v[1]) * (v[2] - v[0])

		var s1 = arrToStr([2]int{v[0], v[1]})
		var s2 = arrToStr([2]int{v[0], v[3]})
		var s3 = arrToStr([2]int{v[2], v[1]})
		var s4 = arrToStr([2]int{v[2], v[3]})

		judge(s1, &set)
		judge(s2, &set)
		judge(s3, &set)
		judge(s4, &set)
	}

	var t1 = arrToStr([2]int{minX, minY})
	var t2 = arrToStr([2]int{minX, MaxY})
	var t3 = arrToStr([2]int{MaxX, minY})
	var t4 = arrToStr([2]int{MaxX, MaxY})

	if len(set) != 4 {
		return false
	}

	if !contain(t1, &set) || !contain(t2, &set) || !contain(t3, &set) || !contain(t4, &set) {
		return false
	}

	return area == (MaxY - minY) * (MaxX - minX)	
}

func contain(s string, m *map[string]void) bool {
	_, ok := (*m)[s]
	return ok
}

func judge(s string, m *map[string]void)  {
	if _, ok := (*m)[s]; ok {
		delete((*m), s)
	} else {
		(*m)[s] = t
	}
}

func arrToStr(arr [2]int) string {
	return strconv.Itoa(arr[0]) + "_" + strconv.Itoa(arr[1])
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

