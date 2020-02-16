package main

import "fmt"

//big结构体的接口类型可以断言为big结构体的内嵌结构体small1或者small2的接口类型，但是反过来不行。
//small1和small2的接口类型不能相互断言。
//Big结构体有接口Big1，Big1有父接口Big2。则Big的类型可以在Big1和Big2中自由切换。
//结构体可以在自己实现的接口和内嵌结构体实现的接口中自由切换
func main() {
	var two Two = &Big{}
	if _, ok := two.(One); ok {
		fmt.Println("two turn to type One")
	}

	var one1 One = &Small1{}
	if _, ok := one1.(Two); !ok {
		fmt.Println("one can't turn to type Two")
	}

	var one2 One = &Small1{}
	if _, ok := one2.(Three); !ok {
		fmt.Println("one can't turn to type Three")
	}

	var big1 Big1 = &Big{}
	if _, ok := big1.(Big2); ok {
		fmt.Println("big1 turn to big2")
	}

	var big2 Big2 = &Big{}
	if _, ok := big2.(Big1); ok {
		fmt.Println("big2 turn to big1")
	}

	var one3 One = &Big{}
	if _, ok := one3.(Big1); ok {
		fmt.Println("成功")
	}
}

type Big struct {
	Small1
	Small2
}

type Big1 interface {
	Ace1() int
}

type Big2 interface {
	Big1
}

func (b *Big) Ace1() int {
	return 1
}

type Small1 struct {
}

type Small2 struct {
}

type One interface {
	Way1() int
}

type Two interface {
	Way2() int
}

type Three interface {
	Way3() int
}

func (s *Small1) Way1() int {
	return 1
}

func (b *Big) Way2() int {
	return 2
}

func (s *Small2) Way3() int {
	return 3
}
