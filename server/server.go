package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP        string
	Port      int
	OnlineMap map[string]*User
	MapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线User
func (s Server) ListenMessage() {
	for true {
		msg := <-s.Message

		s.MapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.MapLock.Unlock()
	}
}

//广播消息的方法
func (s Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.UserAddr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}
func (s *Server) Handler(conn net.Conn) {
	//当前链接的业务...
	//fmt.Println("connected success")
	user := NewUser(conn)
	//用户上先，将用户加入到OnlineMap
	s.MapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.MapLock.Unlock()

	//广播用户上线消息
	s.BroadCast(user, "online")
	//当前Handler阻塞
	select {}
}
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}
func (s *Server) Run() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("Listen failed,err:", err)
	}
	defer listener.Close()
	go s.ListenMessage()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accept failed,err:", err)
			continue
		}
		go s.Handler(conn)
	}
}
