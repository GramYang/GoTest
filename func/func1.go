package main

import "fmt"

func main() {
	//可变函数参数测试
	f1()
}

func f1() {
	op1()
	op1(1)
	op1(1, 2, 3)
}

func op1(i ...int) {
	fmt.Println(len(i))
	for _, v := range i {
		fmt.Println(v)
	}
}
