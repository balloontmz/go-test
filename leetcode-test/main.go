package main 

import (
	"fmt"
)

func main(){
	frequencyForNum := make(map[int]int)
	fmt.Print("不存在的键取出的值为：", frequencyForNum[0])
}