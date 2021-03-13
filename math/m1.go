package main

import (
	"fmt"
	"math"
)

func main() {
	//math包使用
	m1()
}

func m1() {
	a := float32(10) / 3
	fmt.Println(a)                                                                  //3.3333333
	fmt.Println(math.Abs(-3.333))                                                   //3.333
	fmt.Println(math.Float32bits(3.333), math.Float32bits(20.00))                   //1079332831
	fmt.Println(math.Float32frombits(1079332831), math.Float32frombits(1101004800)) //3.333 20
}
