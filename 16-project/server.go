package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("连接建立成功")
}

func (this *Server) Start() {
	//建立一个 socket 监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net Listen err：", err)
	}

	//defer 关闭 socket 监听
	defer listener.Close()

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

func NewServer(ip string, port int) *Server {
	server := &Server{Ip: ip, Port: port}
	return server
}
