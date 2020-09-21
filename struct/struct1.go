package main

import "fmt"

type myStruct struct {
	age  int
	name string
}

func main() {
	//ms1为指针类型
	ms1 := &myStruct{20, "蔡徐坤"}
	//ms2为值类型
	ms2 := myStruct{22, "孙笑川"}
	ms3 := myStruct{}
	modify1(ms2)
	modify2(ms1)
	fmt.Println(ms1)
	fmt.Println(ms2)
	fmt.Println(ms3)
	//输出{0 }，这说明结构体的默认值为其中元素的默认值的结构体集合
	b := &bag{}
	fmt.Println(b) //&{[] map[] <nil> <nil> <nil>}
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
