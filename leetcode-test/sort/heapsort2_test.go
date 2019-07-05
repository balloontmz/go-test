package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort2(t *testing.T)  {
	cases := []int{5, 3, 4, 1, 2}
	HeapSort2(cases)
	fmt.Println("堆排2结果为：", cases)
}