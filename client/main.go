package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
    // conn, err := net.Dial("tcp","139.162.227.61:8000")
    conn, err := net.Dial("tcp",":8000")
    if err !=nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer conn.Close()

    go readMessages(conn)
    
    for {
        // get user input 
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf(">")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("An error occurred reading your input:",err)
            continue
        }
        input = strings.TrimSpace(input) 
        
        _,err =conn.Write([]byte(input))
        if err != nil{
            fmt.Println("Error sending the message")
            continue
        }
    }
}

func readMessages(conn net.Conn){
    for {
        message := make([]byte,1024)
        length, err := conn.Read(message)
        if err != nil {
            fmt.Println("Failed to read message from server:",err)
            os.Exit(1)
        }
        fmt.Println(string(message[:length]))
    }
}
