//Package sort 直接插入排序
/*

*/
package sort

import (
	// "fmt"
)

//DirectInsertionSort 直接插入排序
func DirectInsertionSort(li []int)  {
	l := len(li)
	// 直接插入排序，从索引为 1 开始。
	// 从当前索引开始向前比对，
	for i := 1; i < l; i++ {
		if li[i] < li[i-1] {
			temp := li[i]
			j := i - 1
			for {
				if j < 0 {  // 判断 j 的值需要在取数据的该索引之前。
					break
				}
				if li[j] <= temp {
					break
				}
				
				li[j+1] = li[j]  // 插入操作，是不是只需要进行一次就行了？ -- 不行，插入点之前的数据都需要后移
				j --
			}
			li[j+1] = temp
		}
	}	
}