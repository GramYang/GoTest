package main

import "fmt"

func main() {
	//测试一下[]byte是值类型还是指针类型
	//[]byte是指针，则符合切片的规则
	bt1()
}

func bt1() {
	bs := []byte("我是你爹")
	op1(bs)
	fmt.Println(string(bs)) //d��是你爹
}

func op1(b []byte) {
	b[0] = 100
	fmt.Println(string(b)) //d��是你爹
}
