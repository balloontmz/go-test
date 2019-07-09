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

func TestCoinChange(t *testing.T) {
	cases := []int{25, 21, 10, 5, 1}
	money := 63
	coinUsed := make([]int, money + 1)
	CoinChange(cases, coinUsed, money)
	fmt.Println("硬币找零的结果为：", coinUsed)
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