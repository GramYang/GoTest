package main

import "fmt"

func main() {
	//可变函数参数测试
	f1()
}

func f1() {
	op1()        //0
	op1(1)       //1 1
	op1(1, 2, 3) //3 1 123
}

func op1(i ...int) {
	fmt.Print(len(i), " ")
	if len(i) > 1 {
		fmt.Print(i[0], " ")
	}
	for _, v := range i {
		fmt.Print(v)
	}
	fmt.Println()
}
