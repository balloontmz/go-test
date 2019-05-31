//Hanoi Tower汉诺塔
package main

import (
	"fmt"
)

func hanoi(n int, x, y, z string)  {
	if n == 0 {
		// 等于零的情况下不进行任何操作
	} else {
		hanoi(n - 1, x, z, y)	// x 经过 y 移动到 z
		fmt.Printf("%s -> %s , ", x, y)
		// fmt.Println("%c -> %c , ", x, y)
		hanoi(n - 1, z, y, x) // z 经过 x 移动到 y
	}
}

func main()  {
	hanoi(60, "A", "B", "C")
}