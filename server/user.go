package main

import "net"

type User struct {
	UserAddr string
	Name     string
	C        chan string
	conn     net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		UserAddr: userAddr,
		Name:     userAddr,
		C:        make(chan string),
		conn:     conn,
	}
	//启动监听消息的channel的goroutine

	go user.ListenMessage()
	return user
}

//监听当前userchannel 的方法，一旦有消息，就直接发送给客户端
func (u User) ListenMessage() {
	for true {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
