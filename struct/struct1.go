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
	fmt.Println(ms1)               //&{20 蔡徐坤nmsl}
	fmt.Println(ms2)               //{22 孙笑川}
	fmt.Println(ms3)               //{0 }
	fmt.Println(ms3 == myStruct{}) //true，结构体的值实例是值类型，不能与nil相比，其默认值是其结构体空值
	//输出{0 }，这说明结构体的默认值为其中元素的默认值的结构体集合
	b := &bag{}
	fmt.Println(b) //&{[] map[] <nil> <nil> <nil>}
	//b.B[1]="a"//panic了，说明其指针类型域不会默认初始化
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
