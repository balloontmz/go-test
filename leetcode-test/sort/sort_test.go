package sort

import (
	"fmt"
	"testing"
)

//--------------功能测试--------------------------//
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

/*
硬币找零的动态规划
*/
func TestCoinChange(t *testing.T) {
	cases := []int{25, 21, 10, 5, 1}
	money := 63
	coinUsed := make([]int, money + 1)
	CoinChange(cases, coinUsed, money)
	fmt.Println("硬币找零的结果为：", coinUsed)
}

func TestSimpleSelectionSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	SimpleSelectionSort(cases)
	fmt.Println("简单选择排序结果为：", cases)
}

func TestDirectInsertionSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	DirectInsertionSort(cases)
	fmt.Println("直接插入排序结果为：", cases)
}

func TestMergeSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	MergeSort(cases)
	fmt.Println("直接插入排序结果为：", cases)
}

func TestQuickSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	QuickSort(cases)
	fmt.Println("快速排序结果为：", cases)
}

func TestQuickSort2(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	QuickSort2(cases)
	fmt.Println("快速排序2结果为：", cases)
}

func TestShellSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2, 7, 33, 21, 21, 5, 4}
	ShellSort(cases)
	fmt.Println("希尔排序结果为：", cases)
}

//----------------------------------------------//

//----------------基准测试-----------------------//

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