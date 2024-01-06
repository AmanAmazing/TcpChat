package main

import (
	"fmt"
	"log"
	"net"
	"os"

)


func main() {
	port := ":8000"
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer listner.Close()
	fmt.Println("Server is listening on port: ",port)
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
            continue
		}
		go HandleConn(conn)

	}
}

func HandleConn(conn net.Conn) {
    defer conn.Close()
	fmt.Println("Following IP address connected: ", conn.RemoteAddr())
    buffer := make([]byte,1024)
    for {
        n, err := conn.Read(buffer)
        if err !=nil {
            fmt.Println("Error reading from connection:",err)
            return
        }
        message := string(buffer[:n])
        fmt.Println("Received:",message)
    }
}
