package sort

import (
	// "fmt"
)

/*
这是一个堆排序

堆排序
堆是完全二叉树
其根节点一定是最大或者最小值
堆排序（Heap Sort）就是利用大顶堆或小顶堆的性质进行排序的方法。
堆排序的总体时间复杂度为O(nlogn)。（下面采用大顶堆的方式）
其核心思想是：将待排序的序列构造成一个大顶堆。此时，整个序列的最
大值就是堆的根节点。将它与堆数组的末尾元素交换，然后将剩余的n-1个
序列重新构造成一个大顶堆。反复执行前面的操作，最后获得一个有序序列。
堆排序对原始记录的排序状态不敏感，因此它无论最好、最坏和平均时间复杂度都
是O(nlogn)。在性能上要好于冒泡、简单选择和直接插入算法。
空间复杂度上，只需要一个用于交换的暂存单元。但是由于记录的比较和交换是跳
跃式的，因此，堆排序也是一种不稳定的排序方法。
此外，由于初始构建堆的比较次数较多，堆排序不适合序列个数较少的排序工作。
*/

//HeapSort2 从 python 代码中拷贝过来的
func HeapSort2(li []int)  {
	// fmt.Println("进入排序")
	j := len(li) - 1  // 获取切片长度
	i := (j + 1) / 2
	// 构造一个大顶堆
	for i >= 0 {
		// 从最下面一层有子节点的节点开始，直到顶点 -- 构造顶堆
		heapSort(li, i, j)
		i --
	}
	// 将顶堆顶点交换到最后，然后对顶部节点重新构造大顶堆。循环结束，数据将变成有序
	for j > 0 {
		// 每次循环只需要对顶部的节点进行顶堆运算，其余位置的节点都是已经构造好的。
		swp(li, 0, j)
		heapSort(li, 0, j-1)  // j 以及之后的值已经有序了，不在需要进行排序。
		j --
	}
}

// 构建大顶堆的核心函数
func heapSort(li []int, x, y int)  {
	// 对于构造大顶堆，从下往上构造，依次抛出最大值直到顶点
	// 对于重构大顶堆，从上往下构造，将交换过来的父节点下发至合适的位置
	temp := li[x]
	i := 2 * x   //对于第一次从底部开始构造，传入的 x 为数组长度的一半，那么 i 就是从最底部的或者倒数第二个元素开始。主要看 x 是否整除了数组长度
	for i <= y {
		if i < y && li[i] < li[i+1] { // 选择子节点号，即选择需要弹出值或者放入的分支
			i ++
		} 
		if temp >= li[i] {
			break
		}
		li[x] = li[i] // 较大值放入父节点
		x = i  // 子节点变成父节点 -- 第一次运行初始化时不会再次进入此循环 -- 初始化是多次运行这整个函数
		i *= 2
		li[x] = temp  // 较小值放入被交换的子节点
	}
}