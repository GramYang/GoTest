package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//mutex用法
	test1()
	//第二个测试的输出值恒定为1000，互斥锁指针传递进goroutine，或者限制原生线程数为1
	//test2()
	//稳定，用的atomic，不用互斥锁
	//test3()
	//死锁
	mtt4()
	//测试mutex的特性，当一个goroutine占用锁时，另一个goroutine会等待
	mtt5()
}

type counter struct {
	value int
}

func test1() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()
	for i := 1; i <= 3; i++ {
		wait.Add(1)
		go func(i int, m sync.Mutex) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Lock: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock: ", i)
			mutex.Unlock()
			defer wait.Done()
		}(i, mutex)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait()
}

func test2() {
	//runtime.GOMAXPROCS(1)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	c := &counter{0}
	for i := 0; i < 1000; i++ {
		go func(c *counter, m *sync.Mutex) {
			m.Lock()
			defer m.Unlock()
			c.value++
			wg.Done()
		}(c, &mutex)
	}
	wg.Wait()
	fmt.Println("数数：", c.value)
}

func test3() {
	var wg sync.WaitGroup
	wg.Add(1000)
	var a int32 = 0
	for i := 0; i < 1000; i++ {
		go func(i *int32) {
			for {
				if atomic.CompareAndSwapInt32(i, *i, *i+1) {
					break
				}
			}
			wg.Done()
		}(&a)
	}
	wg.Wait()
	fmt.Println("数数：", a)
}

type Value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func printSum(v1, v2 *Value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()
	time.Sleep(2 * time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()
	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

func mtt4() {
	var a, b Value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

type aaa struct {
	mount int
	m     sync.Mutex
}

func mtt5() {
	a := &aaa{mount: 10}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		a.m.Lock()
		for i := 0; i < a.mount; i++ {
			fmt.Printf("线程1数数：%d\n", i)
			time.Sleep(time.Second)
		}
		a.m.Unlock()
		//a.m.Unlock()//连续解锁两次会panic
		wg.Done()
	}()
	go func() {
		a.m.Lock()
		for i := 0; i < a.mount; i++ {
			fmt.Printf("线程2数数：%d\n", i)
			time.Sleep(time.Second)
		}
		a.m.Unlock()
		wg.Done()
	}()
	wg.Wait()
}
