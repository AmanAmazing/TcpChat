package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
    conn, err := net.Dial("tcp","192.168.0.27:8000")
    if err !=nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer conn.Close()
    
    //buffer := make([]byte,1024)
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
