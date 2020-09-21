package main

import (
	"fmt"
	"reflect"
)

//接口类型可以由interface{}自由转换。
func main() {
	var b IB
	var a = IA(b)
	fmt.Println(reflect.TypeOf(a))
}

type IA interface{}
type IB interface{ foo() int }
