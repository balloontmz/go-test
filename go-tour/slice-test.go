package main

import "golang.org/x/tour/pic"

/**
很有意思的一个切片应用，可以画出各种自定义的图形
*/
func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		temp := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			var num uint8
			num = uint8((i % (j + 1)) / 2)
			temp[j] = num
		}
		ret[i] = temp

	}
	return ret
}

func main() {
	pic.Show(Pic)
}
