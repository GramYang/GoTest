package main

import (
	"fmt"
	"reflect"
)

type func1 func()

func test1() {
	a := func() {}
	b := func1(a)
	fmt.Println(reflect.TypeOf(b))
}

type Ducker interface {
	Quack()
}

type Duck struct {
	name string
}

func (Duck) Quack() {
	fmt.Println("gua gua")
}

type Man struct {
	name string
}

func (Man) Quack() {
	fmt.Println("女王大人")
}

func DoQuack(d Ducker) {
	d.Quack()
}

func test2() {
	d := Duck{"duck"}
	m := Man{"man"}
	DoQuack(d)
	DoQuack(m)
}

func main() {
	//由http包里的HandlerFunc想到的一个小测试
	//test1()
	//go的鸭子类型，这种写法其实java也可以，go只是不用显式写实现interface，看起来像鸭子类型而已
	test2()
}
