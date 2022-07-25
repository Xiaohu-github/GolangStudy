package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip        string           //服务的ip
	Port      int              //服务的ip端口
	OnlineMap map[string]*User //在线用户列表
	mapLock   sync.RWMutex     //同步锁
	Message   chan string      //消息广播的channel
}

//new 一个服务
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听广播消息
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//将msg 发给其他用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	fmt.Println("来自客户端消息：", sendMsg)
	this.Message <- sendMsg
}

//接收消息并广播
func (this *Server) receiveMsg(user *User, conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			this.BroadCast(user, "下线了")
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}

		msg := string(buf[:n-1])
		this.BroadCast(user, msg)
	}
}

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("连接建立成功")
	user := NewUser(conn)
	this.mapLock.Lock()

	this.OnlineMap[user.Name] = user //用户上线,将用户加入到online map 中

	this.mapLock.Unlock()
	this.BroadCast(user, "上线了") //广播当前用户上线了

	this.receiveMsg(user, conn)
}

//启动一个服务端
func (this *Server) Start() {
	//建立一个 socket 监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net Listen err：", err)
	}

	//defer 关闭 socket 监听
	defer listener.Close()

	go this.ListenMessage()

	//循环运行
	for {
		//连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener Accept err：", err)
			continue
		}

		//处理一些操作
		go this.Handler(conn)
	}
}
