package main

import (
	"fmt"
	"sync"
)

//sync.Pool是一个可以存或取的临时对象集合，sync.Pool缓存的期限只是两次gc之间这段时间
//sync.Pool可以安全被多个线程同时使用，保证线程安全
//注意、注意、注意，sync.Pool中保存的任何项都可能随时不做通知的释放掉，所以不适合用于像socket长连接或数据库连接池。
//sync.Pool主要用途是增加临时对象的重用率，减少GC负担。
func main() {
	p := &sync.Pool{New: func() interface{} {
		return 0
	}}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b) //0 1
}
