/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (47.98%)
 * Likes:    2211
 * Dislikes: 180
 * Total Accepted:    404.7K
 * Total Submissions: 828K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 * 
 */
func findKthLargest(nums []int, k int) int {
	index := 1
	switch index {
	case 1:
		return heapSort(nums, k)
	case 2:
		return quickSort(nums, k)
	default:
		bubbleSort(nums)
		return nums[len(nums)-k]
	}
	return 1
}

// 采用 bubbleSort 然后取对应索引的方式
// Your runtime beats 5.17 % of golang submissions
// Your memory usage beats 77.05 % of golang submissions (3.5 MB)
func bubbleSort(li []int)  {
	
	for i := 0; i < len(li) - 1; i++ {  // 排序到倒数第二位
		for j := i+1; j < len(li); j++ {
			if li[i] > li[j] {
				temp := li[i]
				li[i] = li[j]
				li[j] = temp
			}
			
		}
	}
}

// 小顶堆
// ✔ Your runtime beats 99.82 % of golang submissions
// ✔ Your memory usage beats 93.44 % of golang submissions (3.5 MB)
func heapSort(nums []int, k int) int {
	values := nums[len(nums)-k:]
	buildHeap(values)
	for _, v := range nums[:len(nums)-k] {
		if v > values[0] {
			values[0] = v
			adjustHeap(values, 0)	
		}
	}
	return values[0]
}
   
func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- { //////一定得从后往前调整，
		// 切片是引用类型，对其修改会改变原值。
		adjustHeap(values, i)
	}
}
   
func adjustHeap(values []int, pos int) { ///////// 调整pos位置的结点
	node := pos
	length := len(values)
	for node < length {
		var child = 0
		// 从大索引往小索引方向构造顶堆
		if 2*node+2 < length {
			if values[2*node+1] < values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			} ////////选出大子节点
		} else if 2*node+1 < length {
			child = 2*node + 1
		}
		if child > 0 && values[child] < values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}

// 快速排序 -- 每次排完序都能从中间筛选掉一些不再需要排序的数据 
// ✔ Your runtime beats 34.64 % of golang submissions
// ✔ Your memory usage beats 77.05 % of golang submissions (3.5 MB)
func quickSort(li []int, k int) int {
	k = len(li) - k
	low := 0
	high := len(li) - 1
	for low < high {
		pivot := partition(low, high, li)  // 一个关键词索引，所有小于该索引对应的值的值的索引都被置于索引左边，同理
		if pivot == k {
			break
		} else if pivot < k {
			low = pivot + 1
		} else {
			high = pivot - 1
		}
	}
	return li[k]
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
		temp := li[low]
		li[low] = li[high]
		li[high] = temp		
		for low < high {
			if li[low] > pivotValue {
				break
			}
			low++
		}
		temp = li[low]
		li[low] = li[high]
		li[high] = temp
	}
	return low  // pivot 最终所处的索引位置
}
