package main

import (
	"fmt"
	"strings"
)

func main() {
	//对字符串字符下标的测试
	fmt.Println(test1("game#0@nmsl"))
	//对字符串默认值的测试，字符串不会等于nil，只会等于""
	test2()
	//string转rune
	test3()
	//修改字符串中的一个字符
	test4()
	//用切片获取子字符串
	test5()
	//遍历字符串字符
	test6()
}

func test1(s string) string {
	return s[:strings.Index(s, "#")]
}

func test2() {
	var s1 string
	var s2 = ""
	fmt.Println(s1)
	fmt.Println(s2)
	//fmt.Println(s1 == nil)
}

func test3() {
	s := "aaaa"
	t := []rune(s)
	fmt.Println(t)
}

func test4() {
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	fmt.Println(string(c))
}

func test5() {
	str := "hello"
	fmt.Println(str[0:3])
}

func test6() {
	str := "hello"
	for k, v := range str {
		fmt.Printf("%d : %d\n", k, v)
	}
}
