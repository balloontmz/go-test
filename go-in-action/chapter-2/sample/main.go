package main

import (
	"fmt"
	"log"
	"os"

	_ "sample/matchers"
	"sample/search"
)

// main 函数之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
	fmt.Println(1)
}
