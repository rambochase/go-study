package main

import (
	"net"
	"time"
	"log"
	"io"
	"fmt"
)

//nc 127.0.0.1 8080
func main() {
	listener,err := net.Listen("tcp","0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("conn...")
		go handleConn(conn)//不用go就阻塞在这里
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil {
			return  //e.g:
		}
		time.Sleep(time.Second * 1)
	}
}