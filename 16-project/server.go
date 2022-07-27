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

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//存储消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	fmt.Println("来自客户端消息：", sendMsg)
	this.Message <- sendMsg
}

//监听广播消息
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//循环 OnlineMap 将 msg 发给其他用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//处理一些操作
func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	user.Onlie()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			user.Offline()
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}

		msg := string(buf[:n-1])
		user.DoMessage(msg)
	}
}

//启动一个服务端
func (this *Server) Start() {
	//建立监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net Listen err：", err)
	}
	//关闭监听
	defer listener.Close()

	go this.ListenMessage()

	//循环运行
	for {
		conn, err := listener.Accept() //获得连接
		if err != nil {
			fmt.Println("listener Accept err：", err)
			continue
		}
		go this.Handler(conn)
	}

}
