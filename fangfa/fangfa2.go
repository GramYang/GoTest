package main

import "fmt"

//试试接口方法不带参数引用的写法
type a1 interface {
	call(int, string, string)
	speak()
}

type b1 struct{}

func (*b1) call(i int, s1 string, s2 string) {
	fmt.Println("可以实现接口方法call")
}

func (*b1) speak() {}

func (*b1) call1() {}

func (*b1) speak1() {}

func main() {
	b := b1{}
	b.call(1, "1", "1")
}
