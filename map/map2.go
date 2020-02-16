package main

import "fmt"

func main() {
	var m1 map[string]*AAA
	m2 := make(map[string]BBB) //用make来初始化map
	m3 := map[string]BBB{}     //不用make来初始化map
	m2["1"] = &AAA{}
	m1["2"] = &AAA{} //声明的map是nil的不能直接赋值
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}

type AAA struct {
}

type BBB interface {
	Dio()
}

type CCC struct {
	A AAA
	B BBB
}

func (*AAA) Dio() {
	fmt.Println("kuo no Dio da!")
}
