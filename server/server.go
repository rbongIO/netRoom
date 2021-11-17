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
	MapLock   sync.Locker
}

func (s *Server) Handler(conn net.Conn) {
	//当前链接的业务...
	fmt.Println("connected success")
}
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
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
