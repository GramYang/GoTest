package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func DialCustom(network, address string, timeout time.Duration, localIP []byte, localPort int) (net.Conn, error) {
	netAddr := &net.TCPAddr{Port: localPort}
	if len(localIP) != 0 {
		netAddr.IP = localIP
	}
	fmt.Println("netAddr:", netAddr)
	d := net.Dialer{Timeout: timeout, LocalAddr: netAddr}
	return d.Dial(network, address)
}

func main() {
	serverAddr := ":8080"
	var localIP []byte
	localPort := 9001
	conn, err := DialCustom("tcp", serverAddr, time.Second*10, localIP, localPort)
	if err != nil {
		fmt.Println("dial failed:", err)
	}
	defer conn.Close()
	buffer := make([]byte, 512)
	reader := bufio.NewReader(conn)
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}
	fmt.Println("count:", n, "msg:", string(buffer))
	select {}
}
