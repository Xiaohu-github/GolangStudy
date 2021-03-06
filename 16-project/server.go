package main

import (
	"fmt"
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

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("连接建立成功")
	user := NewUser(conn)
	this.mapLock.Lock()

	this.OnlineMap[user.Name] = user //用户上线,将用户加入到online map 中

	this.mapLock.Unlock()
	this.BroadCast(user, "上线了") //广播当前用户上线了
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

//new 一个服务
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	} //new 一个服务
	return server
}
