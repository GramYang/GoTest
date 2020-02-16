package main

import (
	"fmt"
	"reflect"
)

//证明接口引用的类型是指针类型
func main() {
	var x CCC = &BBB{24} //这里的CCC必须是指针类型，不能是值类型
	var m = map[string]CCC{}
	m["nmsl"] = &BBB{24}
	fmt.Println(reflect.TypeOf(x)) //*main.BBB
	fmt.Println(m["nmsl"])
}

type BBB struct {
	age int
}

type CCC interface {
	HowOld() int
}

func (self *BBB) HowOld() int {
	return self.age
}
