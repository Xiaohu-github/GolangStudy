package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	Server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}

	//启动监听当前User channel 的go程
	go user.ListenMessage()

	return user
}

//创建一个go程 User类的方法 ，用来监听当前的 channel,一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

//上线
func (this *User) Onlie() {
	//用户上线,将用户加入到 onlineMap 中
	this.Server.mapLock.Lock()
	this.Server.OnlineMap[this.Name] = this
	this.Server.mapLock.Unlock()
	//广播当前用户上线了
	this.Server.BroadCast(this, "上线了")
}

//下线
func (this *User) Offline() {
	//用户下线,将用户从 onlineMap 中删除
	this.Server.mapLock.Lock()
	delete(this.Server.OnlineMap, this.Name)
	this.Server.mapLock.Unlock()
	//广播当前用户下线
	this.Server.BroadCast(this, "已下线")
}

//给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//消息
func (this *User) DoMessage(msg string) {
	//查询所有在线用户
	if msg == "who" {

		this.Server.mapLock.Lock()
		for _, user := range this.Server.OnlineMap {
			msg := "[" + user.Addr + "]" + user.Name + ":Online...\n"
			this.SendMsg(msg)
		}
		this.Server.mapLock.Unlock()
		//用户名重命名
	} else if len(msg) > 7 && msg[:7] == "rename|" {

		newName := strings.Split(msg, "|")[1]
		_, ok := this.Server.OnlineMap[newName]
		if ok {
			this.SendMsg("The user name already exists !\n")
		} else {
			this.Server.mapLock.Lock()
			delete(this.Server.OnlineMap, this.Name)
			this.Server.OnlineMap[newName] = this
			this.Server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("Your user name updated !\n")
		}

	} else {
		//广播
		this.Server.BroadCast(this, msg)
	}
}
