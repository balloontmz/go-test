package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	// 快排，我擦 一个一个来的那种，基础排序，不是归并排序和希尔排序
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j // 之前忘了这一句，暂时没分析出作用--这一句也是，数组下标的浮动值
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i // 之前忘了这一句，暂时没分析出作用
		}
	}
	values[p] = temp // 也忘了这一句，依旧不知道什么意思
	// 递归
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

// QuickSort 导出的算法
func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
