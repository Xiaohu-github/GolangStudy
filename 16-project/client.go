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
	flag       int //记录当前客户端的模式
}

func NewCLient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error：", err)
		return nil
	}

	client.connet = conn

	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	//从标准输入读取的文本
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}

}

func (clinet *Client) Run() {
	for clinet.flag != 0 {
		for clinet.menu() != true {

		}
		//根据不同的模式处理不同的业务
		switch clinet.flag {

		case 1:
			fmt.Println("公聊模式")
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("更新用户名")
			break
		}

	}
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
	client.Run()
}
