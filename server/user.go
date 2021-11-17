package main

import "net"

type User struct {
	UserAddr string
	UserName string
	Conn     net.Conn
}
