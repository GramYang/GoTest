package main

import "fmt"

func main() {
	//内建函数测试
	bt1()
}

func bt1() {
	//copy测试
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := make([]int, 4)
	copy(arr2, arr1)        //copy只会改变dst不会改变src
	fmt.Println(arr1, arr2) //[1 2 3 4 5] [1 2 3 4]
}
