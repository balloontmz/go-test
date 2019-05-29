package ipc

import (
	"testing"
	// "fmt"
)

type EchoServer struct {
}

func (server *EchoServer) Handler(method, params string) *Response {
	return &Response{"OK", "ECHO: " + method + "~" + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T)  {
	server := NewIpcServer(&EchoServer{})

    // 两个客户端分别调用了 Connect 方法，产生了不同的 session ！！！
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("foo", "From Client1")
	resp2, _ := client2.Call("foo", "From Client2")

	// fmt.Println(resp1.Body)
	// fmt.Println(resp2.Body)
	if resp1.Body != "ECHO: foo~From Client1" || resp2.Body != "ECHO: foo~From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}