package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type netInfo struct {
	Ip     string
	Port   int
	Listen net.Listener
}

//实例化的net.Listener是可以json化的，只不过是一个空的{}
func main() {
	l, err := net.Listen("tcp", ":1213")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	n := netInfo{Ip: "1.1.1.1", Port: 1213, Listen: l}
	data, err2 := json.Marshal(&n)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(string(data)) //{"Ip":"1.1.1.1","Port":1213,"Listen":{}}
}
