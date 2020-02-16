package main

import "fmt"

func main() {
	//数组的最基本使用，go是没有常量数组的，数组可变
	//test1()
	//数组值传递，证明数组是指针类型的
	test2()
}

func test1() {
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}   //指定长度为5，并赋5个初始值
	arr3 := [5]int{1, 2, 3}         //指定长度为5，对前3个元素进行赋值，其他元素为零值
	arr4 := [5]int{4: 1}            //指定长度为5，对第5个元素赋值
	arr5 := [...]int{1, 2, 3, 4, 5} //不指定长度，对数组赋以5个值
	arr6 := [...]int{8: 1}          //不指定长度，对第9个元素（下标为8）赋值1
	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
	//[0 0 0 0 0] [1 2 3 4 5] [1 2 3 0 0] [0 0 0 0 1] [1 2 3 4 5] [0 0 0 0 0 0 0 0 1]
}

func test2() {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("In main(), arr values:", arr)
}

func modify(arr [5]int) {
	arr[0] = 10
	fmt.Println("In modify(), arr values:", arr)
}
