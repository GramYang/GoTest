package main

import "fmt"

func main() {
	fmt.Println(a.varFunc1())
}

var a varA

type varA interface {
	varFunc1() int
	varFunc2() string
}

type varB struct {
}

func (b *varB) varFunc1() int {
	return 1
}

func (b *varB) varFunc2() string {
	return "nmsl"
}
