/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 *
 * https://leetcode.com/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (50.79%)
 * Likes:    601
 * Dislikes: 32
 * Total Accepted:    52.5K
 * Total Submissions: 103.2K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * Implement a MyCalendar class to store your events. A new event can be added
 * if adding the event will not cause a double booking.
 * 
 * Your class will have the method, book(int start, int end). Formally, this
 * represents a booking on the half open interval [start, end), the range of
 * real numbers x such that start <= x < end.
 * 
 * A double booking happens when two events have some non-empty intersection
 * (ie., there is some time that is common to both events.)
 * 
 * For each call to the method MyCalendar.book, return true if the event can be
 * added to the calendar successfully without causing a double booking.
 * Otherwise, return false and do not add the event to the calendar.
 * Your class will be called like this: MyCalendar cal = new MyCalendar();
 * MyCalendar.book(start, end)
 * 
 * Example 1:
 * 
 * 
 * MyCalendar();
 * MyCalendar.book(10, 20); // returns true
 * MyCalendar.book(15, 25); // returns false
 * MyCalendar.book(20, 30); // returns true
 * Explanation: 
 * The first event can be booked.  The second can't because time 15 is already
 * booked by another event.
 * The third event can be booked, as the first event takes every time less than
 * 20, but not including 20.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The number of calls to MyCalendar.book per test case will be at most
 * 1000.
 * In calls to MyCalendar.book(start, end), start and end are integers in the
 * range [0, 10^9].
 * 
 * 
 * 
 * 
 */

// @lc code=start
//我的日历 1
// 108/108 cases passed (104 ms)
// Your runtime beats 36.67 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.8 MB)
type MyCalendar struct {
    data [][2]int
}


func Constructor() MyCalendar {
	var data = make([][2]int, 0)
	return MyCalendar{
		data: data,
	}
}


func (this *MyCalendar) Book(start int, end int) bool {
	var nArr = [2]int{start, end}
	for _, arr := range this.data {
		if verifyIfCross(arr, nArr) {
			return false
		}
	}
	this.data = append(this.data, nArr)
    return true
}

func verifyIfCross(A, B [2]int) bool {
	if B[0] >= A[0] && B[0] < A[1] {
		return true
	}
	if B[1] > A[0] && B[0] < A[1] {
		return true
	}
	if A[0] >= B[0] && A[0] < B[1] {
		return true
	}
	if A[1] > B[0] && A[0] < B[1] {
		return true
	}
	return false
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

