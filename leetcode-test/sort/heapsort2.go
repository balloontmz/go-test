package sort

import (
	"fmt"
)

/*
这是一个堆排序

*/

//HeapSort2 从 python 代码中拷贝过来的
func HeapSort2(li []int)  {
	
}

// 排序交换的原子操作
func swp(li []int, i, j) {
	temp := li[i]
	li[i] = li[j]
	li[j] = temp
}

// 构建大顶堆的核心函数
func heapSort(li []int, x, y int)  {
	temp = li[x]
	
}