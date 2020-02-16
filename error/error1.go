package main

import "fmt"

//go是可以像java那样，函数抛出异常后由外部捕捉处理的
func main() {
	fmt.Println("主线程开始")
	f()
	fmt.Println("主线程结束")
}

func f() {
	defer func() {
		fmt.Println("f()捕捉异常开始")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("f()捕捉异常结束")
	}()
	g()
}

func g() {
	panic("g()主动抛出异常")
}
