package main

import (
	"container/list"
	"fmt"
)

func main() {
	e1 := element{24, "田所浩二"}
	e2 := element{20, "蔡徐坤"}
	e3 := element{26, "孙笑川"}
	test1(e1, e2, e3)
}

type element struct {
	age  int
	name string
}

func test1(es ...element) {
	l := list.New()
	for k := range es {
		l.PushBack(es[k])
	}
	for e := l.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(element)
		if ok {
			fmt.Printf("Age: %d\n", elem.age)
			fmt.Printf("Name: %s\n", elem.name)
		}
	}
}
