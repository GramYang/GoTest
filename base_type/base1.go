package main

import (
	"fmt"
	"reflect"
)

func main() {
	//基本类型转换测试
	//test1()
	//int32负值测试
	test2()
}

func test1() {
	var a int32
	a = 1
	fmt.Println(a == 1) //true
	b := int32(1)
	c := 1
	fmt.Println(reflect.TypeOf(b)) //int32
	fmt.Println(reflect.TypeOf(c)) //int
}

func test2() {
	a := -1
	fmt.Println(int32(a)) //-1
}
