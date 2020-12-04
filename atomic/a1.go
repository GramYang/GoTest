package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//测试atomic.Value使用，不能存nil和不同值类型实例，只能封装了
	atomic_t1()
}

type a struct {
	a1 int
}

type b struct {
	b1 string
}

type c interface {
	op1()
}

func (x *a) op1() {}
func (x *b) op1() {}

func atomic_t1() {
	var v atomic.Value
	fmt.Println(v.Load()) //<nil>
	var a c = &a{}
	var b c = &b{}
	v.Store(a)
	v.Store(b) //报错
}
