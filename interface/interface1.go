package main

import (
	"fmt"
	"reflect"
)

func main() {
	//接口类型可以由interface{}自由转换。
	//t1()
	//证明接口引用的类型是指针类型
	//t2()
	//鸭子类型，go的鸭子类型，这种写法其实java也可以，go只是不用显式写实现interface，看起来像鸭子类型而已
	//t3()
	//interface和nil的区别
	t4()
}

type IA interface{}
type IB interface{ foo() int }

func t1() {
	var b IB
	var a = IA(b)
	fmt.Println(reflect.TypeOf(a)) //<nil>
}

type BBB struct {
	age int
}

type CCC interface {
	HowOld() int
}

func (self *BBB) HowOld() int {
	return self.age
}

func t2() {
	var x CCC = &BBB{24} //这里的CCC必须是指针类型，不能是值类型
	var m = map[string]CCC{}
	m["nmsl"] = &BBB{24}
	fmt.Println(reflect.TypeOf(x)) //*main.BBB
	fmt.Println(m["nmsl"])
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

func t3() {
	d := Duck{"duck"}
	m := Man{"man"}
	DoQuack(d)
	DoQuack(m)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

func t4() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	//为什么会这样？因为interface的==只会判断动态类型，不会判断动态值。那么如何判断动态值呢？
	fmt.Println(IsNil(b)) //true
}
