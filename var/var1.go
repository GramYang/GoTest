package main

import "fmt"

func main() {
	//值语义和指针语义的测试，即还原forrange bug的原理
	vt1()
}

func vt1() {
	var v1 = 1
	var v2 = &v1
	var v3 = v2
	fmt.Println(v1, *v2, *v3) //1 1 1
	v1 = 10                   //修改一个值，就会影响另外两个指针
	fmt.Println(v1, *v2, *v3) //10 10 10
}
