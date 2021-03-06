package main

import (
	"GoTest/etcd/srv_find"
	"fmt"
	"github.com/prometheus/common/log"
	"time"
)

func main() {
	m, err := srv_find.NewMaster([]string{
		"http://127.0.0.1:2379",
	}, "services/")
	if err != nil {
		log.Fatal(err)
	}
	for {
		for k, v := range m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
