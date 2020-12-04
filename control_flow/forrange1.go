package main

import "fmt"

func main() {
	//test1()
	//test2()
	//go中的for，判断条件不满足即终止遍历
	t3()
}

func test1() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	for t := range nums { //slice单值遍历出的是key
		fmt.Println("slice single: ", t)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		k += "nmsl"
		fmt.Println(k, v)
	}
	for v := range kvs {
		fmt.Println(v)
	} //map单值遍历出来的是key，和slice一样
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func test2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for value := range ch {
		fmt.Print(value)
	}
}

func t3() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a) && a[i] != 3; i++ {
		fmt.Println(a[i])
	} //输出1 2
}
