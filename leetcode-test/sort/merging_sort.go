//Package sort test
/*
从简单的开始，暂时搁置。
归并排序
*/
package sort

import (
	"fmt"
)

//MergeSort 归并排序 -- 
func MergeSort(li []int)  {
	fmt.Println("进入排序")
	msort(li, li, 0, len(li)-1)
}

func msort(liL, liR []int, x, y int)  {
	temp := make([]int, len(liL))
	if x == y {
		liR[x] = liL[x]
	} else {
		m := (x + y) / 2
		// fmt.Println("当前递归的中间值为：", m, "x, y 的值为：", x, y)
		// 将数组赋值给 temp
		msort(liL, temp, x, m)
		msort(liL, temp, m + 1, y)  // 此处 m 记得 +1 ，m 这个索引已经归于上一个递归函数了
		// 合并当前递归的 temp 临时变量到父级递归的临时变量 -- 合并完成后，当前递归的临时变量将失去意义
		// 运行这个合并函数之前，m 前后的序列均已经有序了。这次合并的意义类似于合并两个有序的数组。
		merge(temp, liR, x, m, y)
		fmt.Println("排序的结果为：", temp)
	}
}

func merge(liL, liR []int, l, m, r int)  {
	j := m + 1
	k := l
	// 此次合并直到某个索引超限，代表已经有一个数组添加完成，另一个则需要全部添加到 LiR 数组中
	for {
		if l > m {
			break
		}
		if j > r {
			break
		}
		// i 对应左边的序列，j 对应右边的的序列。分别已经有序了，但是整体并不是有序。
		if liL[l] < liL[j] {
			liR[k] = liL[l]
			l ++
		} else {
			liR[k] = liL[j]
			j ++
		}
		k++
	}

	// 注意边界问题
	if l <= m {
		for index := 0; index <= m - l; index++ {
			liR[k + index] = liL[l + index]
		}
	}

	if j <= r {
		for index := 0; index <= r - j; index++ {
			liR[k + index] = liL[j + index]
		}
	}

}