package pack1

//首字母小写的接口是没办法被实现的
type test1 interface {
	aaaa(int) int
	bbbb(int) int
}

//接口里不能有首字母小写的私有方法，不然不能被实现
type Test2 interface {
	Cccc(int) int
	Dddd(int) int
	//eeee(int) int
}
