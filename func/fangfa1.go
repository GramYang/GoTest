package main

import "fmt"

func main() {
	ff1()
	ff2()
}

func ff1() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	b := B{}
	b.Print()
	fmt.Println(b.Name)
	c := A{"nmsl"}
	fmt.Println(c)
	d := C{24}
	(&d).print1()
	d.print1()
	d.print2()
}

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	age int
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}

func (c C) print1() {
	fmt.Println(c.age)
}

func (c *C) print2() {
	fmt.Println(c.age)
}

func ff2() {
	b := b1{}
	b.call(1, "1", "1")
}

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
