package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//该方法没有设置随机因子，不能生成随机序列
	test1()
	//添加了随机因子
	test2()
}

func test1() {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}

func test2() {
	rand.Seed(time.Now().UnixNano())
	sixah := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(sixah), func(i, j int) {
		sixah[i], sixah[j] = sixah[j], sixah[i]
	})
	fmt.Println(sixah)
}
