package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup的使用测试
func main() {
	//这个例子：wg的值传递和引用传递没有区别
	testWG1()
}

func testWG1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(n int, wg1 *sync.WaitGroup) {
			wg.Add(1)
			fmt.Println(n)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("exit")
}
