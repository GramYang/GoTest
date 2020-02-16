package main

import (
	"fmt"
	"time"
)

//父线程提前退出，子线程也会被中断
func main() {
	go counting(10, 30)
	counting(0, 10)
	time.Sleep(time.Second * 15) //不加这个go里面的函数就执行不完
}

func counting(start int, end int) {
	for i := start; i < end; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
