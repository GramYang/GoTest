package main

import (
	"fmt"
	"reflect"
)

//reflect.TypeOf专门测试，返回的全是main.AA
func main() {
	//这里是显式类型转换，*AA是类型，nil是值，得出的结果就是AA的指针，指向结果是nil
	a := reflect.TypeOf((*AA)(nil)).Elem()
	fmt.Println(a)
	fmt.Printf("%s\n", a)

	a1 := reflect.TypeOf(AA{"孙笑川"})
	fmt.Println(a1)
	fmt.Printf("%+v\n", &AA{"孙笑川"})

	var b AA
	a2 := reflect.TypeOf(b)
	fmt.Println(a2)

	var c *AA
	a3 := reflect.TypeOf(c).Elem()
	fmt.Println(a3)

	var d BB = (*AA)(nil)
	a4 := reflect.TypeOf(d).Elem()
	fmt.Println(a4)

	a5 := reflect.New(a).Interface()
	fmt.Println(reflect.TypeOf(a5))
}

type AA struct {
	s1 string
}

type BB interface {
	age() int
}

func (a *AA) age() int {
	return 1
}
