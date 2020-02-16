package main

import (
	"fmt"
	"reflect"
)

func main() {
	//基本使用
	//test1()
	//测试map是指针类型
	//test2()
	//test3()
	//结构体里面map初始化测试，结构体中的map必须单独make
	test4()
}

func test1() {
	//var m1 map[int]string
	//m1[1] = "nmsl" //声明的map是nil，不能赋值
	m2 := make(map[int]string)
	m2[1] = "nmsl"
	fmt.Println(m2)
	fmt.Println(reflect.TypeOf(m2))
}

func test2() {
	m := make(map[int]int)
	mdMap1(m)
	fmt.Println(m) //map[1:100 2:200]
}

func test3() {
	m := make(map[int]int)
	mdMap2(m)
	fmt.Println(m) //map[]
}

func mdMap1(m map[int]int) {
	m[1] = 100
	m[2] = 200
}

func mdMap2(m map[int]int) {
	m = make(map[int]int)
	m[1] = 100
	m[2] = 200
}

type wrapmap struct {
	M1 map[int]string
}

func test4() {
	wm := &wrapmap{}
	//wm := &wrapmap{M1:make(map[int]string)}
	wm.M1[1] = "nsml" //panic: assignment to entry in nil map
	fmt.Println(wm.M1[1])
}
