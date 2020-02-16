package main

//测试"神奇的写法"。
// 简单来说就是解决了结构体的接口变量不能调用其内嵌结构体方法的问题。
func main() {
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
