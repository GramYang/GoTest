package main

import "fmt"

func main() {
	//test1()
	//test2()
	test3()
}

func test1() {
	var count int
	var sum int
	fmt.Println("请输入切片长度：")
	_, _ = fmt.Scanf("%d", &count)
	s := make([]int, count)
	for i := 0; i < len(s); i++ {
		fmt.Printf("请输入第%d个数：\n", i+1)
		_, _ = fmt.Scanf("%d", &s[i])
		//_, _ = fmt.Scanf("%d\n", &s[i])
	}
	for j := 0; j < len(s); j++ {
		sum += s[j]
	}
	fmt.Println("切片求和：", sum)
}

func test2() {
	var float1 = 1231321345.123123
	fmt.Printf("%9.1f\n", float1)
	var string1 = "1"
	fmt.Printf("%5s\n", string1)
}

func test3() {
	fmt.Println(fmt.Sprintf("%s#%d@%s", "game", 0, "dev"))
}
