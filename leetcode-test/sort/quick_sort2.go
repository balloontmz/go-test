package sort

import (
	"fmt"
)

//QuickSort2 快排
func QuickSort2(li []int) {
	fmt.Println("快速排序 2")
	qSort2(0, len(li)-1, li)
}

func qSort2(begin, end int, li []int) {
	if begin >= end {
		return
	}
	i := begin
	key := li[begin]
	for j := begin + 1; j <= end; j ++ {
		if li[j] < key {
			// begin 所在的元素只在遍历完成后进行交换
			i ++
			swp(li, i, j)
		}
	}
	swp(li, begin, i)
	// i 所在元素已经有序，不需要再放进序列中排序
	qSort2(begin, i-1, li)
	qSort2(i+1, end, li)
}