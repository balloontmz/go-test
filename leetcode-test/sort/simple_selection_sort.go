//Package sort 简单选择排序
/*

*/
package sort

//SimpleSelectionSort 简单选择排序
func SimpleSelectionSort(li []int)  {
	l := len(li)
	for i := 0; i < l; i++ {
		mininum = i
		for j := i+1; j <= l; j++ {
			if li[mininum] > li[j] {
				mininum = j
			}
		}
	}
}