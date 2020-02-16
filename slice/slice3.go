package main

import (
	"fmt"
	"sort"
)

func main() {
	//slice实现动态数组
	//test5()
	//切片边界测试
	//test6()
	//切片排序，需要实现sort包的Interface接口
	//test7()
	//测试结构体中的切片初始化，其中的切片必须make
	//test8()
	//切片与可变长参数的转换，可变长参数能赋值给切片引用
	test9()
}

func test5() {
	var a []int                 //这里用生命代替make，不然会多上len个0
	fmt.Println(a == nil)       //true
	fmt.Println(len(a), cap(a)) // 0 0
	for i := 0; i < 20; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a)) // 长度为1+1，容量为1-2-4-8-16-32，在长度等于容量后容量才会扩容
	}
	fmt.Println(a)
	a = nil //使用完之后可以清空为声明状态
	fmt.Println(a)
	for i := 0; i < 20; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
}

func test6() {
	var a []int
	for i := 0; i < 20; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	b := a[:17]
	fmt.Println(b) //[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	c := a[17:]
	fmt.Println(c) //[17 18 19]
}

type person struct {
	Age int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func test7() {
	persons := personSlice{
		person{
			Age: 1,
		},
		person{
			Age: 5,
		},
		person{
			Age: 2,
		},
	}
	sort.Sort(persons)
	fmt.Printf("after sort:%+v", persons)
}

type wrapslice struct {
	S []int
}

func test8() {
	//ws := &wrapslice{S:make([]int,10)}
	ws := new(wrapslice)
	ws.S[0] = 1
}

func test9() {
	x:=[]int{1,2,3,4,5,6,7,8,9,10}
	foo(x...)
}

func foo(arr ...int){
	s1:=arr[0:5]
	fmt.Println(s1) //[1 2 3 4 5]
	fmt.Println(len(s1), cap(s1)) //5 10
	fmt.Println(arr[0:5]) //[1 2 3 4 5]
}