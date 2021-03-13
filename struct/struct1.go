package main

import "fmt"

type myStruct struct {
	age  int
	name string
}

func main() {
	//基本使用
	s1()
	//多态
	//s2()
	//神奇的写法，巧妙的使用了go的伪鸭子类型实现了内嵌结构体转换
	//s3()
}

func modify1(m myStruct) {
	m.name = m.name + "nmsl"
}

//只有指针类型才能修改
func modify2(m *myStruct) {
	m.name = m.name + "nmsl"
}

type bag struct {
	A []int
	B map[int]string
	C chan int
	D interface{}
	E func(int) int
}

func s1() {
	//ms1为指针类型
	ms1 := &myStruct{20, "蔡徐坤"}
	//ms2为值类型
	ms2 := myStruct{22, "孙笑川"}
	ms3 := myStruct{}
	var ms4 myStruct
	modify1(ms2)
	modify2(ms1)
	fmt.Println(ms1)                                    //&{20 蔡徐坤nmsl}
	fmt.Println(ms2)                                    //{22 孙笑川}
	fmt.Println(ms3)                                    //{0 }
	fmt.Println(ms3 == myStruct{}, &ms3 == &myStruct{}) //true false，结构体的值实例是值类型，不能与nil相比，其默认值是其结构体空值
	fmt.Println(ms2 == myStruct{})                      //false
	fmt.Println(ms4, ms4 == myStruct{})                 //{0 } true
	b := &bag{}
	fmt.Println(b) //&{[] map[] <nil> <nil> <nil>}
	//b.B[1]="a"//panic了，说明其指针类型域不会默认初始化
}

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

func s2() {
	f := father{}
	s := son{}
	s.speak()
	f.speak() //没有定义father方法时调用的是son的方法，视为重写
}

type a1 struct{}

type a2 struct {
	a1
}

type b1 interface {
	fun1()
}

type b2 interface {
	fun2()
}

func (*a1) fun1() {}

func (*a2) fun2() {}

func s3() {
	//本身结构体是可以调用其内嵌结构体方法的
	//x := a2{}
	//x.fun1()
	//x.fun2()
	//y := a1{}
	//y.fun1()
	var x1 b1 = &a1{}
	x1.fun1() //但是结构变量就不行
	x1.(interface {
		fun2()
	}).fun2()
}
