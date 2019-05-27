package ipc

/*
这个包封装了 session 的生成。
提供了 Server 接口用于定义处理请求的方法。

*/

import (
	"encoding/json"
	"fmt"
)

type (
	Request struct {
		Method string `json:"method"`
		Params string `json:"params"`
	}

	Response struct {
		Code string `json:"code"`
		Body string `json:"body"`
	}

	Server interface {
		Name() string
		Handler(method, params string) *Response
	}

	IpcServer struct {
		Server
	}
)

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	// 请求没有传递进来吧？ 还是此方法仅用于定义连接
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" { // 关闭连接
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
				return
			}

			resp := server.Handler(req.Method, req.Params)

			b, err := json.Marshal(resp)

			c <- string(b) // 返回结果
		}
	}(session)

	fmt.Print("新 session 创建成功")
	return session
}
