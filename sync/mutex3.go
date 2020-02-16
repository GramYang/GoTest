package main

import (
	"fmt"
	"sync"
	"time"
)

type aaa struct {
	mount int
	m     sync.Mutex
}

//测试mutex的特性，当一个goroutine占用锁时，另一个goroutine会等待
func main() {
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
