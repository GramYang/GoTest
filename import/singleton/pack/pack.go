package pack

import "fmt"

var (
	B = get()
)

type A struct {
	State bool
}

func get() *A {
	fmt.Println("执行一次初始化")
	return &A{true}
}
