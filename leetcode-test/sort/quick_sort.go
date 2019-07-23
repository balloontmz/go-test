//Package sort 此文件用于快速排序徐
/*
快速排序基于冒泡排序
复杂度为 o(nlogn)
核心思想为将待排序记录分隔成独立的两部分，其中一部分的值均小于另一部分，然后对其分别进行排序 -- 跟归并排序的流程正好相反
*/
package sort

import (
	"fmt"
)

//QuickSort 快速排序
func QuickSort(li []int)  {
	fmt.Println("开始快速排序")
	qSort(0, len(li)-1, li)
}

func qSort(low, high int, li []int)  {
	if low < high {
		pivot := partition(low, high, li)  // 一个关键词索引，所有小于该索引对应的值的值的索引都被置于索引左边，同理
		qSort(low, pivot, li)
		qSort(pivot+1, high, li)
	}
}

func partition(low, high int, li []int) int {
	// 快速排序核心代码
	pivotValue := li[low]
	for low < high {
		for low < high {
			if li[high] < pivotValue {
				break
			}
			high--
		}
		swp(li, low, high)
		for low < high {
			if li[low] > pivotValue {
				break
			}
			low++
		}
		swp(li, low, high) // 此处可能两种情况 索引相等或者 low 值大于 high 值
	}
	return low  // pivot 最终所处的索引位置
}