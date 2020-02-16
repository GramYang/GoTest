package main

import (
	"fmt"
	"reflect"
)

//结构体的反射
func main() {
	//test3()
	//test4()
}

type Type1 struct {
	s1, s2, s3 string
}

func (t Type1) String() string {
	return t.s1 + "-" + t.s2 + "-" + t.s3
}

var secret interface{} = Type1{"Ada", "Go", "Oberon"}

type T struct {
	A int
	B string
}

func test3() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)
	//main.Type1
	//struct
	//	Field 0: Ada
	//	Field 1: Go
	//	Field 2: Oberon
	//[Ada-Go-Oberon]
}

func test4() {
	t := T{23, "skidoo"}
	fmt.Println(t)
	//因为Elem()，所以ValueOf里面必须传入指针
	s := reflect.ValueOf(&t).Elem()
	fmt.Println(s)
	typeOfT := s.Type()
	fmt.Println(typeOfT)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
	//{23 skidoo}
	//{23 skidoo}
	//main.T
	//0: A int = 23
	//1: B string = skidoo
	//t is now {77 Sunset Strip}
}
