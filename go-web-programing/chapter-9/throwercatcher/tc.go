package main

import (
	"fmt"
	"time"
)

func thrower (c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >> ", i)
	}
}

func catcher(c chan int)  {
	for i := 0; i < 5; i++ {
		num := <- c
		fmt.Println("Caught <<", num)
	}
}

func main()  {
	c := make(chan int)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
	/*
	一下存在 caught 在 threw 之前的，是因为当 写入 chan 就被读出来了并且调度先执行了 caught 的语句。
	Caught << 0
	Threw >>  0
	Threw >>  1
	Caught << 1
	Caught << 2
	Threw >>  2
	Threw >>  3
	Caught << 3
	Caught << 4
	Threw >>  4
	*/
}