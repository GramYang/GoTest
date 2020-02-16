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

func main() {
	//由http包里的HandlerFunc想到的一个小测试
	test1()
}
