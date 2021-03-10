package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//读锁无效
	//readlock()
	//写锁:1 write end结束之后，2才能reading；2 read end结束之后，3 才能writing
	//writelock()
	rw1()
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

type bag struct {
	l   sync.RWMutex
	m   map[int]int
	len int
}

func (b *bag) read(w *sync.WaitGroup, key int) {
	b.l.RLock()
	v, ok := b.m[key]
	b.l.RUnlock()
	fmt.Println(v, ok)
	w.Done()
}

func (b *bag) write(w *sync.WaitGroup, key, value int) {
	b.l.Lock()
	b.m[key] = value
	b.len++
	b.l.Unlock()
	w.Done()
}

func rw1() {
	var w sync.WaitGroup
	b := &bag{m: make(map[int]int)}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go b.write(&w, i, i+10)
	}
	w.Add(1)
	go b.read(&w, 1)
	w.Wait()
	fmt.Println(b.m, b.len)
}
