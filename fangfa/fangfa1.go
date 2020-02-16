package main

import "fmt"

func main() {
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
