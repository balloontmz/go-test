package main 

import (
	"fmt"
	// "time"
)

// func callerA(c chan string)  {
// 	// time.Sleep(1 * time.Microsecond)
// 	c <- "Hello A!"
// }

// func callerB(c chan string)  {
// 	// time.Sleep(1 * time.Microsecond)
// 	c <- "Hello B!"
// }

// func main() {
// 	a, b := make(chan string), make(chan string)
// 	go callerA(a)
// 	go callerB(b)
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Microsecond) // 防止 goroutine 尚未开始执行循环就结束了
// 		select {
// 		case msg := <-a:
// 			fmt.Printf("%s from A\n", msg)
// 		case msg := <-b:
// 			fmt.Printf("%s from B\n", msg)
// 		default:
// 			fmt.Println("Default") // 防止通道不存在值时的阻塞
// 		}
// 	}
// }

func callerA(c chan string)  {
	// time.Sleep(1 * time.Microsecond)
	c <- "Hello A!"
	close(c)
}

func callerB(c chan string)  {
	// time.Sleep(1 * time.Microsecond)
	c <- "Hello B!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a: // 采用双值接收，第二个参数判定通道是否关闭
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)	
			}
		}
	}
}