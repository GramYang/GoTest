package main

import "fmt"

func main() {
	//测试go里面的并发读，看来并发读是不会出错的，参考gin里面对Engine实例的并发读
	t1()
}

func t1() {
	b := &Bag{a: 123, b: "abc"}
	for i := 0; i < 50; i++ {
		go b.op1()
	}
}

type Bag struct {
	a int64
	b string
}

func (b *Bag) op1() {
	fmt.Println(b.a, b.b)
}
