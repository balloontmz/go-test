//Package sort 简单选择排序
/*

*/
package sort

import (
	"fmt"
)

//SimpleSelectionSort 简单选择排序  -- 对应堆排，每次选出一个最大或者最小的推出。
func SimpleSelectionSort(li []int)  {
	l := len(li)
	for i := 0; i < l - 1; i++ {
		mininum := i
		for j := i+1; j < l; j++ {
			if li[mininum] > li[j] {  // 先比较，保存数组的索引。
				mininum = j
			}	
		}
		if i != mininum {
			fmt.Println("最小索引值为：", mininum)
			swp(li, i, mininum)
		}
	}
}