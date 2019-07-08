package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	BubbleSort(cases, 1)
	fmt.Println("冒泡排序结果为结果为：", cases)
}

func TestAdvanceBubbleSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	BubbleSort(cases, 2)
	fmt.Println("冒泡排序结果为结果为：", cases)
}

func  BenchmarkBubbleSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
        cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
		BubbleSort(cases, 1)
		// fmt.Println("冒泡排序结果为结果为：", cases)
    }
	
}

func  BenchmarkHeapSort2Sort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
        cases := []int{5, 3, 4, 1, 2}
		HeapSort2(cases)
		// fmt.Println("冒泡排序结果为结果为：", cases)
    }
	
}
func  BenchmarkAdvanceBubbleSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
        cases := []int{5, 3, 4, 1, 2}
		BubbleSort(cases, 2)
		// fmt.Println("冒泡排序结果为结果为：", cases)
    }
}