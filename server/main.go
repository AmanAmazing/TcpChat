package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Client struct {
	conn net.Conn
	ch   chan<- string
}

var (
	clients    = make(map[Client]bool)
	register   = make(chan Client)
	unregister = make(chan Client)
	broadcast  = make(chan string)
	mutex      = &sync.Mutex{}
)

func main() {
	port := ":8000"
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer listner.Close()
	fmt.Println("Server is listening on port: ", port)
    go handleChannels()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
        ch := make(chan string)
        client := Client{conn,ch}
        register <- client
		go HandleConn(client)
	}
}

func HandleConn(client Client){
	defer client.conn.Close()
	fmt.Println("Following IP address connected: ", client.conn.RemoteAddr())
	buffer := make([]byte, 1024)
	for {
		n, err := client.conn.Read(buffer)
		if err != nil {
            unregister <- client
            break
		}
		message := string(buffer[:n])
        broadcast <- message
	}
}

func handleChannels(){
    for {
        select{
        case client := <-register:
            mutex.Lock()
            clients[client] = true
            mutex.Unlock()
        case client := <-unregister:
            mutex.Lock()
            if _,ok := clients[client]; ok {
                close(client.ch)
                delete(clients,client)
            }
            mutex.Unlock()
        case message := <- broadcast: 
            for client := range clients{
                client.ch <- message
            }
        }
    }
}
