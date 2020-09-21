package main

import (
	"fmt"
	"strings"
)

func main() {
	stringsT1()
}

func stringsT1() {
	s1 := "我是你哥哥"
	s2 := "我是你哥哥_我们两个都是你妈的儿子"
	fmt.Println(strings.Split(s1, "_")) //[我是你哥哥]
	fmt.Println(strings.Split(s2, "_")) //[我是你哥哥 我们两个都是你妈的儿子]
	s3 := "nmsl"
	fmt.Println(strings.ToUpper(s3[0:1]) + strings.ToLower(s3[1:])) //Nmsl
}
