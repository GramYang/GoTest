package main

import "fmt"

func main() {
	//continue+label
	//t1()
	//goto+label
	//t2()
}

func t1() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1 //每一轮都不输出j=4和j=5
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func t2() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE //label必须在定义后面
}
