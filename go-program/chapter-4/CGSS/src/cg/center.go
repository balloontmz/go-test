//Package cg center 中央处理器
//TODO: 时间问题此包函数未全部完成，后期再补。此文件重点是读写锁！！！
package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"cgss/ipc"
)

var _ ipc.Server = &CenterServer{} // 确认实现了Server接口

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

// 这里 `...` 中的文本叫做 tag，它只是普通字符串，你也可以理解为对结构体字段的注释，但它对反射机制可见
// 很多编解码器都会利用反射机制的这个能力做编解码，这里是 JSON 编解码器需要

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []int // 此结构体似乎还未实现？？？
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()

	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, player)

	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	return nil
}

// 返回值中定义了返回参数名字，则函数中不再需要重新声明？？？
func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()

	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online")
	}
	return
}

func (server *CenterServer) broadcast(params string) error {
	return nil
}

func (server *CenterServer) Handler(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	default:
		return &ipc.Response{Code: "404", Body: method + ": " + params}
	}
	return &ipc.Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
