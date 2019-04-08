package bubblesort

//Bubblesort 冒泡排序实现算法
func Bubblesort(values []int) {
	flag := true

	for i := 0; i < len(values); i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				// 判定是否进行了交换，如果没有，则代表其有序
				flag = false
			}
		}
		// 如果一次遍历完没有进行过交换，则代表有序!!!
		if flag == true {
			break
		}
	}
}
