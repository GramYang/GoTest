package main

import (
	"fmt"
	"reflect"
)

func main() {
	//基本使用
	test1()
	//使用set来修改值
	//t2()
}

type Wrap struct {
	A int
	B string
	C []int
	D map[int]string
}

func test1() {
	//基本类型
	var x = 3.4
	fmt.Println(reflect.TypeOf(x))               //float64
	fmt.Println(reflect.ValueOf(x))              //3.4
	fmt.Println(reflect.ValueOf(x).Type())       //float64
	fmt.Println(reflect.ValueOf(x).Kind())       //float64
	fmt.Println(reflect.ValueOf(x).Float())      //3.4
	fmt.Println(reflect.ValueOf(x).Interface())  //3.4，这是将Value转换为Interface{}的关键方法
	fmt.Println(reflect.Zero(reflect.TypeOf(x))) //0
	//结构体
	w := Wrap{A: 100, B: "nmsl", C: []int{1, 2, 3, 4}, D: map[int]string{1: "蔡徐坤", 2: "孙笑川"}}
	fmt.Println(reflect.TypeOf(w))              //main.wrap
	fmt.Println(reflect.ValueOf(w))             //{100 nmsl [1 2 3 4] map[1:蔡徐坤 2:孙笑川]}
	fmt.Println(reflect.ValueOf(w).Kind())      //struct
	fmt.Println(reflect.ValueOf(w).Interface()) //{100 nmsl [1 2 3 4] map[1:蔡徐坤 2:孙笑川]}
	z := reflect.Zero(reflect.TypeOf(w))
	fmt.Println(z, z.Interface().(Wrap).C == nil, z.Interface().(Wrap).D == nil) //{0  [] map[]} true true
	//封装类型
	s := []int{1, 2, 3, 4, 5}
	m := map[int]string{1: "蔡徐坤", 2: "孙笑川"}
	fmt.Println(reflect.TypeOf(s), reflect.ValueOf(s), reflect.ValueOf(s).Kind(), reflect.Zero(reflect.TypeOf(s)))
	//[]int [1 2 3 4 5] slice []
	fmt.Println(reflect.ValueOf(s).Interface() == reflect.Zero(reflect.TypeOf(s))) //false
	//本来类型是interface{}的map或slice是不能比较的，一比较就会panic，但是这里可以比较，因为reflect.Zero(reflect.TypeOf(s))是nil
	fmt.Println(reflect.TypeOf(m), reflect.ValueOf(m), reflect.ValueOf(m).Kind(), reflect.Zero(reflect.TypeOf(m)))
	//map[int]string map[1:蔡徐坤 2:孙笑川] map map[]
	fmt.Println(reflect.ValueOf(m).Interface() == reflect.Zero(reflect.TypeOf(m))) //false
	//map[int]string map[1:蔡徐坤 2:孙笑川] map map[]
	//检查bean里面是否有域未被初始化，序列化大型bean时使用
	w1 := Wrap{A: 200}
	w1v := reflect.ValueOf(w1)
	wv := reflect.ValueOf(w)
	for i := 0; i < w1v.NumField(); i++ {
		if w1v.Field(i).Kind() == reflect.Map || w1v.Field(i).Kind() == reflect.Slice {
			if w1v.Field(i).Interface() == nil {
				fmt.Println("w1有未初始化域")
			}
		} else if w1v.Field(i).Interface() == reflect.Zero(w1v.Field(i).Type()).Interface() {
			fmt.Println("w1有未初始化域")
		}
	}
	for i := 0; i < wv.NumField(); i++ {
		if wv.Field(i).Kind() == reflect.Map || wv.Field(i).Kind() == reflect.Slice {
			if wv.Field(i).Interface() == nil {
				fmt.Println("w有未初始化域")
			}
		} else if wv.Field(i).Interface() == reflect.Zero(wv.Field(i).Type()).Interface() {
			fmt.Println("w有未初始化域")
		}
	}
	//反射判断一个实例是否被初始化
	w2 := &Wrap{}
	fmt.Println(reflect.Indirect(reflect.ValueOf(w2)) == reflect.Zero(reflect.TypeOf(Wrap{}))) //false
	fmt.Println(reflect.Indirect(reflect.ValueOf(w2)), reflect.Zero(reflect.TypeOf(Wrap{})))
	//{0  [] map[]} {0  [] map[]}
	//Value是不能相互比较的，只能Interface{}转换后再进行比较
}

type wrap1 struct {
	a int
}

func t2() {
	var a = 1
	var b = &wrap1{a: 100}
	v1 := reflect.Indirect(reflect.ValueOf(&a))
	v2 := reflect.Indirect(reflect.ValueOf(b))
	fmt.Println(v1, v2) //1 {100}
	v1.Set(reflect.ValueOf(100))
	v2.Set(reflect.ValueOf(wrap1{a: 200}))
	fmt.Println(a, b) //100 &{200}
}
