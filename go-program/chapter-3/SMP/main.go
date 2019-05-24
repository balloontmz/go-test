package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
 
	"smp/mlib"
	"smp/mp"
)

var lib *mlib.MusicManager
var id int = 1
var ctrl, signal chan int

func handlerLibCommands(tokens []string){
	switch tokens[1] {
		case "list":
			for i := 0; i < lib.Len(); i++ {
				e, _ := lib.Get(i)
				fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
			}
		case "add": {
			fmt.Println("命令长度为;", len(tokens))
			if len(tokens) == 6 {
				
				id++
				lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><artist><source><type>")
			}
		}
		case "remove":
			if len(tokens) == 3 {
				//TODO: 此处代码存在问题
				lib.Remove(1)	// 此处是根据 id 删除，不合理 (int)(tokens[2])
			} else {
				fmt.Println("USAGE: lib remove <name>")
			}
		default:
			fmt.Println("Unrecognized lib command:", tokens[1])
		
	}
}

func handlePlayCommand(tokens []string)  {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return 
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The Music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)	// , ctrl, signal 此处功能并不完整
}

func main()  {
	fmt.Println(`
			Enter ...
	`)
	lib = mlib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for{
		fmt.Println("Enter command-> ")

		rawline, _, _ := r.ReadLine()

		line := string(rawline)

		if line == "q" || line =="e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handlerLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("未知命令:", tokens[0])
		}
	}
}