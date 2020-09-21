package main

import (
	"fmt"
	"time"
)

//goroutine也是一个闭包
func main() {
	a := &ABC{1}
	go a.counting1()
	//go a.counting2()
	time.Sleep(time.Second * 2)
	fmt.Printf("在父goroutine中：%d\n", a.count)
}

type ABC struct {
	count int
}

func (self ABC) counting1() {
	self.count += 2
	fmt.Printf("在子goroutine中：%d\n", self.count)
}

func (self *ABC) counting2() {
	self.count += 2
	fmt.Printf("在子goroutine中：%d\n", self.count)
}
