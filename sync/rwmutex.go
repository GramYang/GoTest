package main

import (
	"sync"
	"time"
)

func main() {
	//读锁无效
	//readlock()
	//写锁:1 write end结束之后，2才能reading；2 read end结束之后，3 才能writing
	writelock()
}

func readlock() {
	var m sync.RWMutex
	go read(&m, 1)
	go read(&m, 2)
	time.Sleep(2 * time.Second)
}

func read(m *sync.RWMutex, i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "read end")
}

func writelock() {
	var m sync.RWMutex
	go write(&m, 1)
	go read(&m, 2)
	go write(&m, 3)
	time.Sleep(4 * time.Second)
}

func write(m *sync.RWMutex, i int) {
	println(i, "write start")
	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "write end")
}
