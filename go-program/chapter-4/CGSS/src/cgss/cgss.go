package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"

	"cgss/cg"
	"cgss/ipc"
)

var centerClient  *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}

	return nil
}

func Help(args []string) int {
	// 多行文本？
	fmt.Println(`
	Commands:
		login <username><level><exp>
		logout <username>
		send <message>
		listplayer
		quit(q)
		help(h)
	`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}

	centerClient.RemovePlayer(args[1])

	return 0
}

func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}

	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("<level> 必须为整数")
		return 0
	}

	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("<exp> 必须为整数")
		return 0
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp

	err = centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("添加用户失败", err)
	}

	return 0
}

func ListPlayer(args []string) int {
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("失败：", err)
	} else {
		for i, v := range ps {
			fmt.Println(i + 1, ":", v)
		}
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")

	err := centerClient.Broadcast(message)
	if err != nil {
		fmt.Println("失败：", err)
	}
	return 0
}

//GetCommandHandlers 注释
// map[string](func (args []string) int) 键是字符串，值是函数的映射。看到得不多，打个标签
func GetCommandHandlers() map[string](func (args []string) int) {
	return map[string](func (args []string) int) {
		"help": Help,
		"h": Help,
		"quit": Quit,
		"q": Quit,
		"login": Login,
		"logout": Logout,
		"listplayer": ListPlayer,
		"send": Send,
	}
}

func main()  {
	fmt.Println("Casual Game Server Solution")

	startCenterService()

	Help(nil)

	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for{
		fmt.Println("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok {  // ok 能判断是否取到值，相关查看官方文档
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("未知的命令", tokens[0])
		}
	}
}