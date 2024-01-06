package main

import (
	"fmt"
	"log"
	"net"
)

type Room struct {
	Name         string
	Participants []string
	Full         bool
}

var Chatroom = Room{
	Name:         "testing",
	Participants: make([]string, 5),
	Full:         false,
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}
		go HandleConn(conn)

	}
}

func HandleConn(conn net.Conn) {
	fmt.Println("Following IP address connected: ", conn.RemoteAddr())
    Chatroom.Participants = append(Chatroom.Participants, conn.RemoteAddr().String())
	n, err := conn.Write([]byte("this is not encrypted!"))
	if err != nil {
        // need to remove the connection from the chatroom participants pool  
        fmt.Println("Connection closed for: ",conn.RemoteAddr().String())
        conn.Close()
	}
	fmt.Println(n)
}
