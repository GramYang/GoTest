package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//输出常量
	//test1()
	//原生线程调度
	//test2() //有序的
	//test3() //乱序的
	test4() //乱序的
}

func test1() {
	fmt.Println("系统: ", runtime.GOOS)
	fmt.Println("CPU几核：", runtime.NumCPU())
	fmt.Println("Go的根目录", runtime.GOROOT())
}

func test2() {
	runtime.GOMAXPROCS(1)
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
		if i == 4 {
			runtime.Gosched() //让当前goroutine让出CPU
		}
	}
	<-exit
}

func test3() {
	runtime.GOMAXPROCS(2)
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)

		if i == 4 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}

func test4() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
