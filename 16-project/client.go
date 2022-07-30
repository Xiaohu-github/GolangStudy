package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	name       string
	connet     net.Conn
}

func NewCLient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error：", err)
		return nil
	}

	client.connet = conn

	return client
}

func main() {
	client := NewCLient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>>连接服务器失败...")
		return
	}
	fmt.Println(">>>>>>>>连接成功...")

	//启动客户端的业务
	select {}
}
