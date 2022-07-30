package main

import (
	"flag"
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

var serverIp string
var serverPort int

//使用时： .\client.exe -ip 127.0.0.1 -port 8888
//初始方法init 固定写法，用于执行之前进行一个初始化工作
func init() {
	//新增指令
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewCLient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>连接服务器失败...")
		return
	}
	fmt.Println(">>>>>>>>连接成功...")

	//启动客户端的业务
	select {}
}
