package main

import (
	"fmt"
	"reflect"
)

func main() {
	//接口类型可以由interface{}自由转换。
	//t1()
	//证明接口引用的类型是指针类型，除了interface{}之外
	//t2()
	//鸭子类型，go的鸭子类型，这种写法其实java也可以，go只是不用显式写实现interface，看起来像鸭子类型而已
	//t3()
	//interface和nil的区别
	//t4()
	//interface是实现在type上，严格来说已经没有指针和值的区别了，定义的type都是一个新的值类型type，不管你是什么
	t5()
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
	fmt.Println(m["nmsl"])         //&{24}
	var x1 interface{} = BBB{25}   //但是interface{}不同，interface{}可以是任何类型
	var x2 interface{} = x
	fmt.Println(reflect.TypeOf(x1), reflect.TypeOf(x2)) //main.BBB *main.BBB
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
	var bbb interface{} = (*BBB)(nil)
	var ccc CCC = (*BBB)(nil)
	fmt.Println(bbb == ccc)                                //true
	fmt.Println(reflect.TypeOf(bbb), reflect.ValueOf(bbb)) //*main.BBB <nil>
	bbb1 := reflect.New(reflect.TypeOf(bbb).Elem())        //(*BBB)(nil)的使用，利用其Type生成一个&BBB{}
	fmt.Println(bbb1)                                      //&{0}
}

type AAAA interface {
	op1() int
}

type BBBB int
type CCCC []int
type DDDD map[int]int
type EEEE chan int

func (*BBBB) op1() int {
	return 2
}
func (*CCCC) op1() int {
	return 3
}
func (*DDDD) op1() int {
	return 4
}
func (*EEEE) op1() int {
	return 5
}

func t5() {
	var t1 BBBB = 10                      //必须这么写
	var x1 AAAA = &t1                     //只能用指针类型
	fmt.Println(reflect.TypeOf(x1))       //*main.BBBB
	var t2 CCCC = []int{1, 2, 3}          //同上
	var x2 AAAA = &t2                     //同上
	fmt.Println(reflect.TypeOf(x2))       //*main.CCCC
	var t3 DDDD = map[int]int{1: 1, 2: 2} //同上
	var x3 AAAA = &t3                     //同上
	fmt.Println(reflect.TypeOf(x3))       //*main.DDDD
	var t4 EEEE = make(chan int)          //同上
	var x4 AAAA = &t4                     //同上
	fmt.Println(reflect.TypeOf(x4))       //*main.EEEE
}
