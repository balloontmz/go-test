package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("第一个切片为：", slice1)
	slice2 := slice1[1:3]
	fmt.Println("第二个切片为：", slice2)
	slice1 = append(slice1, 99) // 此处添加一个元素，导致 slice1 不再引用之前的底层数组，后面的值互不影响
	slice2[1] = 10
	fmt.Println("第一个切片此时为：", slice1, "对应的第二个切片此时为：", slice2)
	slice2 = append(slice2, 10)
	fmt.Println("第一个切片此时为：", slice1, "对应的第二个切片此时为：", slice2)
}
