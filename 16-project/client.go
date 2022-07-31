package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
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

func (client *Client) DealResponse() {
	//如果client.connet有数据，就直接copy到stdout标准输出上
	//永久阻塞=>for{}
	io.Copy(os.Stdout, client.connet)
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

func (clinet *Client) updateName() bool {
	fmt.Println(">>>>>>请输入用户名")
	fmt.Scanln(&clinet.Name)

	sendMsg := "rename|" + clinet.Name + "\n"
	_, err := clinet.connet.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

func (clinet *Client) Run() {
	for clinet.flag != 0 {
		for clinet.menu() != true {

		}
		//根据不同的模式处理不同的业务
		switch clinet.flag {

		case 1:
			fmt.Println("公聊模式")
			var chatMsg string
			fmt.Println(">>>>>>请输入聊天内容,exit退出")
			fmt.Scanln(&chatMsg)
			for chatMsg != "exit" {
				if len(chatMsg) != 0 {
					sendMsg := chatMsg + "\n"
					_, err := clinet.connet.Write([]byte(sendMsg))
					if err != nil {
						fmt.Println("connet Write error:", err)
						break
					}
				}
				chatMsg = ""
				fmt.Println(">>>>>>请输入聊天内容,exit退出")
				fmt.Scanln(&chatMsg)
			}
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("更新用户名")
			clinet.updateName()
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

	go client.DealResponse()
	//启动客户端的业务
	client.Run()
}
