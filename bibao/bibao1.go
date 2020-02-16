package main

import (
	"fmt"
	"runtime"
	"sync"
)

//闭包测试
//本来是一个闭包的测试，但是引出了另一个问题，
// 就是golang的原生线程调度而出现的并发问题，这里牵涉到runtime包的使用
func main() {
	//test1()
	//输出10个值，大部分是10，还有几个是随机数
	//test2()
	//乱序输出0-9
	//test3()
	//一个通道只能阻塞一个goroutine，和通道的缓冲无关
	test4()
}

//这里go中直接使用外部的waitgroup构成闭包，因为waitgroup的值传递和引用传递都是一样的，所以使用闭包没有问题
func test1() {
	var s = []string{
		"first",
		"second",
		"third",
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(s))
	for _, item := range s {
		go func(i string) {
			fmt.Println(i)
			//fmt.Println(item)
			waitGroup.Done()
		}(item)
	}
	waitGroup.Wait()
}

func test2() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("测试1", i)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func test3() {
	runtime.GOMAXPROCS(1) //加了这个后，乱序问题解决了
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for j := 0; j < 10; j++ {
		go func(a int) {
			fmt.Println("测试2", a)
			waitGroup.Done()
		}(j)
	}
	waitGroup.Wait()
}

func test4() {
	runtime.GOMAXPROCS(1) //加了这个后，乱序问题解决了
	c := make(chan bool, 10)
	for j := 0; j < 10; j++ {
		go func(a int) {
			fmt.Println("测试2", a)
			c <- true
		}(j)
	}
	<-c
}
