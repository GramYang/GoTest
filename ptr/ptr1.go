package main

import "fmt"

//https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
func main() {
	//golang里面只有值传递
	test1()
	test2()
	test3()
}

func test1() {
	i := 10
	ip := &i
	fmt.Println("i的内存地址是：", &i)
	fmt.Println("ip的值：", ip)
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify1(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify1(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}

func test2() {
	persons := make(map[string]int)
	persons["张三"] = 19
	mp := &persons
	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modify2(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func modify2(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}

func test3() {
	ages := []int{6, 6, 6}
	fmt.Printf("原始slice的内存地址是%p\n", ages)
	modify3(ages)
	fmt.Println(ages)
}

func modify3(ages []int) {
	fmt.Printf("函数里接收到slice的内存地址是%p\n", ages)
	ages[0] = 1
}
