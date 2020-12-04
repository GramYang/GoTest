package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
)

func main() {
	//基本api测试
	l, _ := lru.New(5)
	fmt.Println(l.Add(1, "a")) //false
	v, ok := l.Get(2)
	fmt.Println(v, ok) //<nil> false
	v, ok = l.Get(1)
	fmt.Println(v, ok) //a true
	v, ok = l.Get(1)
	fmt.Println(v, ok)           //a true
	fmt.Println(l.Contains(1))   //true
	fmt.Println(l.Contains(100)) //false
	//缓存值修改
	m := map[int]string{1: "a", 2: "b"}
	l.Add(10, m)
	v, ok = l.Get(10)
	v.(map[int]string)[2] += "nsml"
	v1, _ := l.Get(10)
	fmt.Println(v1)            //map[1:a 2:bnsml]
	l.Add(20, &wrap{a: "蔡徐坤"}) //只有指针类型结构体引用才能修改，如果结构体的域是一个引用，就不需要指针类型
	v2, ok := l.Get(20)
	v2.(*wrap).a += "nmsl"
	v3, _ := l.Get(20)
	fmt.Println(v3) //&{蔡徐坤nmsl}
}

type wrap struct {
	a string
}
