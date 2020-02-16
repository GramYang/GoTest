package pack2

//结构体实现接口和其的包访问性无关
type aaa struct {
}

type AAA struct {
}

func (*aaa) Cccc(a int) int {
	return a + 1
}

func (*aaa) Dddd(a int) int {
	return a + 1
}

//func (*aaa) eeee(a int) int {
//	return a + 1
//}
