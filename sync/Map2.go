package main

import (
	"fmt"
	"sync"
)

var myMap sync.Map

func main() {
	myMap.Store("1", []string{"hi"})
	myMap.Store(1, 11111)
	if val, ok := myMap.Load("1"); ok {
		fmt.Println("查", val)
	}
	myMap.Store("1", 2222222)
	if val, ok := myMap.Load("1"); ok {
		fmt.Println("改", val)
	}
	if v, ok := myMap.LoadOrStore("22", "33333333"); ok {
		fmt.Println("LoadOrStore", v)
	}
	myMap.Delete(1)
	f := func(key, value interface{}) bool {
		fmt.Println("遍历", key, value)
		return true
	}
	myMap.Range(f)
}
