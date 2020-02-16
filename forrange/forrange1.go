package main

import "fmt"

func main() {
	//test1()
	test2()
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
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
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
