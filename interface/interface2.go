package main

import (
	"fmt"
	"reflect"
)

//接口类型可以自由转换，这是golang的特点或者说是缺陷
func main() {
	var b IB
	var a = IA(b)
	fmt.Println(reflect.TypeOf(a))
}

type IA interface{}
type IB interface{ foo() int }
