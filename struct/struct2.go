package main

import (
	"fmt"
)

type father struct {
	son
}

type son struct {
}

func (s *son) speak() {
	fmt.Println("我是儿子")
}

func (f *father) speak() {
	fmt.Println("我是你爹")
}

//结构体内嵌实现父子结构体，golang可以重写但是不能重载
func main() {
	f := father{}
	s := son{}
	s.speak()
	f.speak() //没有定义father方法时调用的是son的方法，视为重写
}
