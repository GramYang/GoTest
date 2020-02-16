package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//测试对通道输入nil的结果：
	//如果通道的数据类型是int之类的，输入nil会编译不通过。
	//而如果数据类型是默认值为nil之类的，那就正常获取nil。
	//test1()
	//一个奇怪的姿势，在select中向通道写入，
	// 这么写有什么好处？见channel2
	test2()
}

func test1() {
	ch := make(chan error)
	go func(ch chan error) {
		ch <- errors.New("first")
		time.Sleep(time.Second * 2)
		ch <- errors.New("second")
		time.Sleep(time.Second * 2)
		ch <- nil
	}(ch)
	go func(ch chan error) {
		for {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println(i)
				} else {
					fmt.Println("ch已关闭")
				}
			case err := <-ch:
				fmt.Println(err)
			}
		}

	}(ch)
	time.Sleep(time.Second * 10)
}

func test2() {
	ch := make(chan int)

	go func(ch chan int) {
		select {
		case ch <- 1:
			fmt.Println("send 1")
		}
	}(ch)

	v := <-ch
	fmt.Println("recv: ", v)
}
