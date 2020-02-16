package main

import (
	"fmt"
)

//golang可以重写，不能重载，重载可以靠向方法内传递interface{}类型的参数来实现
func main() {
	test1()
	//test2()
}

func test1() {
	c := C1{}
	fmt.Println(c.Say())
	fmt.Println(c.B1.Say())
	fmt.Println(c.A1.Say())
}

func test2() {
	c := &C1{}
	fmt.Println(c.Say())
	fmt.Println(c.B1.Say())
	fmt.Println(c.A1.Say())
}

type A1 struct {
}

func (a A1) Say() string {
	return "a"
}

type B1 struct {
	A1
}

func (b B1) Say() string {
	return "b"
}

//func (b B1) Say(s string) string {
//	return s
//}

type C1 struct {
	B1
}

func (c C1) Say() string {
	return "c"
}
