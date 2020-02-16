package main

import (
	"fmt"
	"time"
)

//time.AfterFunc会和主线程逻辑交替执行
func main() {
	time.AfterFunc(time.Second*5, func() {
		for j := 10; j < 20; j++ {
			fmt.Println(j)
			time.Sleep(time.Second)
		}
	})
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 10) //不加这一句AfterFunc的回调会执行不完
}
