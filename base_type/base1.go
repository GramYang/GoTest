package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//基本类型转换测试
	//test1()
	//整数相除获得浮点数的坑，小数点后位数会被省略掉
	t2()
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

func t2() {
	//计算后的结果类型必须要显式指定，不然就和计算参数类型一致
	var a interface{} = 10 / 11
	var a1 float32 = 10 / 11
	var a2 = float32(10 / 3)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a1), reflect.TypeOf(a2)) //int float32 float32
	fmt.Println(a, a1, a2, a2 == 3)                                        //0 0 3 true
	//不同类型的计算参数不能计算
	var a3 int8 = 10
	var a4 int64 = 3
	var a5 = a3 / int8(a4)
	fmt.Println(reflect.TypeOf(a5), a5) //int8 3
	//必须这么写才能保证浮点值小数点后位数不被省略
	var a6 = float32(30) / float32(100)
	fmt.Println(a6, a6 == 0) //0.3 false
	//控制浮点数小数后多少位的唯一方法
	v1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0.2223), 64)
	fmt.Println(v1) //0.22 这里必须是64位
	v2, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", 0.255), 64)
	fmt.Println(v2) //0.3
}
