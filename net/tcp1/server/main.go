package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Printf("net.ResolveTCPAddr failed:%s", addr)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("listen %s failed:%s", addr, err)
		return
	} else {
		log.Println("rpc listening", addr)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept error :", err)
			continue
		}
		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	buffer := []byte("You are welcome. I'm server.")
	n, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Write error:", err)
	}
	fmt.Println("send:", n)
	fmt.Println("connection end")
}
