/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var p, q int
	var temp []int
    for index := 0;index < m + n; index ++ {
		if p == len(nums1) || nums1[p] == 0 {
			temp = append(temp, nums2[q:]...)
			break
		} else if q == len(nums2) || nums2[q] == 0 {
			temp = append(temp, nums1[p:]...)
			break
		}
		if nums1[p] > nums2[q] {
			temp = append(temp, nums2[q])
			q++ 
		} else {
			temp = append(temp, nums1[p])
			p++ 
		}
	}
	return temp
}

