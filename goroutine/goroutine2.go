package main

import (
	"fmt"
	"time"
)

//goroutine中运行方法
//看来和a是值类型还是指针类型无关，和方法的接受者是值类型还是指针类型有关
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
