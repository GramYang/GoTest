package main

import (
	"fmt"
	"reflect"
)

func main() {
	//基本使用
	test1()
	//测试map是指针类型
	//test2()
	//test3()
	//结构体里面map初始化测试，结构体中的map必须单独make
	//test4()
	//遍历修改map里的值
	//t5()
	//map和*map函数传参的区别，同slice，不过map没有append
	//t6()
}

func test1() {
	//var m1 map[int]string
	//m1[1] = "nmsl" //声明的map是nil，不能赋值
	m2 := make(map[int]string)
	m2[1] = "nmsl"
	fmt.Println(m2)
	fmt.Println(reflect.TypeOf(m2))
	v, ok := m2[2]
	fmt.Println(v == "", ok)                    //true false
	m3 := map[string]string{"a": "b", "c": "d"} //另一种初始化方法
	fmt.Println(m3["a"], m3["c"])
	fmt.Println(m3["z"] == "") //true
	var m4 map[int]string
	fmt.Println(m4)        //map[]
	fmt.Println(len(m4))   //0
	fmt.Println(m4 == nil) //true
	m5 := make(map[int]string)
	fmt.Println(m5)        //map[]
	fmt.Println(len(m5))   //0
	fmt.Println(m5 == nil) //false
	m6 := make(map[string][]int)
	m6["nmsl"] = []int{1, 2, 3, 4}
	m6["12"] = make([]int, 5)
	m6["34"] = []int{}
	fmt.Println(m6)              //map[12:[0 0 0 0 0] 34:[] nmsl:[1 2 3 4]]
	fmt.Println(m6["34"] == nil) //false
	m7 := make(map[string]interface{})
	m7["a"] = "a"
	m7["b"] = 1
	fmt.Println(m7["c"] == nil) //true
	m8 := map[int]string{}
	fmt.Println(m8 == nil) //false
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

func t5() {
	m1 := map[string]string{"1": "1", "2": "2", "3": "3"}
	//for k,v:=range m1{
	//	if k=="1"{
	//		k="a"
	//	}
	//	if v=="2"{
	//		v="b"
	//	}
	//}
	//上面的写法是不能修改map的，下面的写法可以修改，但是map的key顺序会变
	//因为程序执行过程中map的长度会变化，为了map值的正确，go语言不允许直接修改map中的值类型结构。
	for k, v := range m1 {
		if k == "1" {
			delete(m1, k)
			m1["a"] = v
		}
		if v == "2" {
			m1[k] = "b"
		}
	}
	fmt.Println(m1)
}

func t6() {
	a := map[int]int{1: 1, 2: 2}
	b := map[int]int{3: 3, 4: 4}
	op1(a, &b)
	fmt.Println(a, b) //map[1:101 2:2] map[3:103 4:4]
}

func op1(a map[int]int, b *map[int]int) {
	a[1] += 100
	(*b)[3] += 100
}
