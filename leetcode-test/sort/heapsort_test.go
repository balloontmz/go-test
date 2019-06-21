package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2}
	HeapSort(cases)
	fmt.Println("堆排结果为：", cases)
}