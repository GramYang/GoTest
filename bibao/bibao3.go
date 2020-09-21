package main

import "sync"

func main() {
	//测试golang中的闭包持有的是值还是指针
	//测试结果：闭包持有的是指针，因此不会拷贝值
	//而const对象是禁止的，完美搭配
	//simpleTest1()
	//使用闭包和mutex并发修改一个外部捕获引用的值，goroutine没有join，只能用waitgroup来阻塞主协程
	simpleTest2()
}

func simpleTest1() {
	o := object{20}
	func() {
		o.A += 20
		println(o.A)
	}()
	println(o.A)
}

func simpleTest2() {
	o := object{20}
	var m sync.Mutex
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			m.Lock()
			o.A += 1
			println(o.A)
			m.Unlock()
			w.Done()
		}()
	}
	w.Wait()
	println(o.A)
}

type object struct {
	A int
}
