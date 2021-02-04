package main

import "fmt"

//这是一个纯正的闭包测试
func main() {
	//闭包：带函数中的函数带return
	//testBibao1()
	//闭包引用，说明了闭包持有的是指针而不是值
	testBibao2()
	//循环闭包
	//testBibao3()
	//循环闭包，输出3个3
	//testBibao4()
	//针对上面例子的修改
	//testBibao5()
}

func testBibao1() {
	i := incr()
	//一个闭包，三次调用，闭包是可以修改外部引用变量（基本类型）
	fmt.Println(i()) //1
	fmt.Println(i()) //2
	fmt.Println(i()) //3
	//三个闭包
	fmt.Println(incr()()) //1
	fmt.Println(incr()()) //1
	fmt.Println(incr()()) //1
}

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func testBibao2() {
	x := 1
	func() {
		fmt.Println(x) //1
		x++
		fmt.Println(x)  //2
		fmt.Println(&x) //0xc0000140a0
	}()
	fmt.Println(&x) //0xc0000140a0
	x = 2
	func() {
		fmt.Println(x) //2
	}()
	x = 3
}

func testBibao3() {
	var f func()
	for i := 0; i < 3; i++ {
		f = func() {
			fmt.Println(i)
		}
	}
	f() //3
}

func testBibao4() {
	var funcSlice []func()
	for i := 0; i < 3; i++ {
		funcSlice = append(funcSlice, func() {
			println(i)
		})
	}
	for _, v := range funcSlice {
		v()
	}
}

func testBibao5() {
	var funcSlice []func()
	for i := 0; i < 3; i++ {
		func(i int) {
			funcSlice = append(funcSlice, func() {
				println(i)
			})
		}(i)
	}
	for _, v := range funcSlice {
		v()
	}
}
