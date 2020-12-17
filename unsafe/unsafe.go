package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//unsafe使用测试
	unt1()
}

type bag struct {
	a bool
	b uint8
	c int
	d string
}

func unt1() {
	var s1 float64 = 1.15
	s2 := *(*uint64)(unsafe.Pointer(&s1)) //强行转换类型
	fmt.Println(reflect.TypeOf(s2))       //uint64
	b := bag{true, 30, 1, "hello"}
	fmt.Println(unsafe.Sizeof(b))     //32
	fmt.Println(unsafe.Alignof(b.b))  //1
	fmt.Println(unsafe.Offsetof(b.d)) //16
	fmt.Printf("%p\n", &b)            //0xc0000044a0
	p := unsafe.Pointer(&b)
	fmt.Println(p) //0xc0000044a0
	//uintptr是golang的内置类型，是能存储指针的整型
	s3 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.c)))
	*s3 = 100
	fmt.Println(b.c) //100
}
