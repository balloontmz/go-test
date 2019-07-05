package sort

// 排序交换的原子操作
func swp(li []int, i, j int) {
	temp := li[i]
	li[i] = li[j]
	li[j] = temp
}