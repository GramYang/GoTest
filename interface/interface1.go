package main

//测试接口的实现方法
func main() {

}

type A interface {
	AA1() int
	AA2() int
	AA3() string
}

//B必须实现A的所有方法才算是实现了A
type B struct{}

//而结构体内嵌A，没有实现任何A的方法也算是实现了A
type C struct {
	A
	B string
}

func (b *B) AA1() int {
	return 1
}

func (b *B) AA2() int {
	return 2
}

func (b *B) AA3() string {
	return "蔡徐坤"
}
