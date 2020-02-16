package main

import (
	"fmt"
	"sync"
)

var m sync.Map

type userInfo struct {
	Name string
	Age  int
}

func main() {
	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load(1)
	fmt.Println(vv, ok) //<nil> false

	vv, ok = m.Load("2")
	fmt.Println(vv, ok) //<nil> false

	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //oneone true

	m.Store(1, "oneoneone") //这里会同时存在两个：1 oneoneone，1的类型不同

	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChanPly"
	user1.Age = 24
	map1["user1"] = user1

	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")
	for k, v := range mapValue.(interface{}).(map[string]userInfo) {
		fmt.Println(k, v)
	}
}
