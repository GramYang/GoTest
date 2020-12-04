package main

import "fmt"

//适配器模式，用结构体封装结构体并用新方法调用其方法
func main() {
	oa := new(oldAdaptee)
	adapter := new(Adapter)
	adapter.adaptee = oa
	adapter.Process()
}

type Target interface {
	Process()
}

type Adaptee interface {
	Foo()
	Bar()
}

type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Process() {
	fmt.Println("在Adapter中执行process()")
	a.adaptee.Bar()
	a.adaptee.Foo()
}

type oldAdaptee struct{}

func (o *oldAdaptee) Foo() {
	fmt.Println("在旧接口中执行foo()")
}

func (o *oldAdaptee) Bar() {
	fmt.Println("在旧的接口中执行bar()")
}
