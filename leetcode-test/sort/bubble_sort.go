//Package sort bubble
/*
冒泡排序：时间复杂度 o(n^2)
核心思想是：两两比较相邻的关键字，如果反序则交换，直到没有反序记录为止。
其实现细节可以不同：
1.最简单排序实现：bubble_sort_simple
2.冒泡排序：bubble_sort
3.改进的冒泡排序：bubble_sort_advance
*/
package sort

import (
	"fmt"
)

//BubbleSort 冒泡排序
func BubbleSort(li []int, index int)  {
	switch index {
	case 1:
		sort(li)
	case 2:
		advanceSort(li)
	default:
		fmt.Println("未定义的算法")
	}
}

func sort(li []int)  {
	for i := 0; i < len(li) - 1; i++ {  // 排序到倒数第二位
		for j := i+1; j < len(li); j++ {
			if li[i] > li[j] {
				swp(li, i, j)  // 包内函数	
			}
			
		}
	}
}

func advanceSort(li []int)  {
	for i := 0; i < len(li) - 1; i++ {  // 排序到倒数第二位
		flag := false    
		for j := i+1; j < len(li); j++ {
			if li[i] > li[j] {
				swp(li, i, j)  // 包内函数	
				flag = true
			}
			
		}
		if flag == false {// 如果一轮没进行交换，则退出
			break
		}
	}
}