package main

import "fmt"

func main() {
	//测试切片的赋值，以及切片的清空
	test1()
	//测试切片的函数传递，切片是指针类型，只不过append是特例
	//append在没有扩容的情况下，会在原来的[]int拼接上值然后返回，
	//虽然返回的指针是不一样的，append并没有生产一个新的[]int。
	//golang就不存在两个变量地址相同的情况
	//test2()
	//panic: runtime error: slice bounds out of range [:1] with capacity 0
	//test3()
	//切片是指针类型
	//test4()
	//测试得出的结果，当一个函数的参数是变长参数，传入切片时必须带上...
	//如果是interface{}时，看似不带...可以编译通过，其实是错误的。相当于用interface{}指代[]interface{}
	//t5()
	//[]int和*[]int函数传参的区别，可以看出：下标寻址都是一样的，但是append只有*[]int会有效
	//t6()
	//测试append创建新切片，但是底层的数组是一样的。测试结果：假的，根本不需要append
	//t7()
}

func test1() {
	a := make([]int, 10)
	a = append(a, 1)
	b := a[:]
	a = append(a, 2)
	fmt.Println(a) //1 2
	fmt.Println(b) //1
	//测试清空切片
	c := make([]int, 10, 10)
	c = c[0:0]
	//c=[]int{} //这样输出的len和cap都是0
	fmt.Println(c)               //[]
	fmt.Println("len: ", len(c)) //0
	fmt.Println("cap: ", cap(c)) //10
	fmt.Println(c == nil)        //false
	var d []int
	fmt.Println(d)        //[]
	fmt.Println(len(d))   //0
	fmt.Println(cap(d))   //0
	fmt.Println(d == nil) //true
	c1 := make([]int, 10, 20)
	fmt.Println(len(c1), cap(c1)) //10 20
	c1[0], c1[1], c1[2], c1[3], c1[4] = 1, 2, 3, 4, 5
	c1 = c1[:len(c1):len(c1)]         //这种写法可以让切片的长度等于容量
	fmt.Println(c1, len(c1), cap(c1)) //[1 2 3 4 5 0 0 0 0 0] 10 10
	c2 := []int{1, 2, 3, 4}
	fmt.Println(len(c2), cap(c2)) //4 4
	c2 = c2[1:2:3]
	fmt.Println(c2, len(c2), cap(c2)) //[2] 1 2
	//比较切片是否相等，切片是不能比较的，数组才能比较
	//c3:=[]int{1,2,3}
	//c4:=[]int{1,2,3,4}
	//c5:=[]int{1,2,3}
	//fmt.Println(c3==c4,c3==c5)
}

func test2() {
	a := make([]int, 10, 20)
	fmt.Printf("%p\n", &a[0])
	add(a)
	var b []int
	fmt.Println(a)      //[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(cap(a)) //10
	fmt.Println(b)      //[]
	add(b)
	fmt.Println(b) //[]
	addPtr(&a)
	addPtr(&b)
	fmt.Println(a) //[0 0 0 0 0 0 0 0 0 0 9]
	var as int32 = 10
	fmt.Println(a[as]) //切片下标用int32也是可以的
	fmt.Println(b)     //[9]
}

func add(x []int) {
	fmt.Printf("%p\n", &x) //变了
	fmt.Println(x)         //[0 0 0 0 0 0 0 0 0 0]
	y := append(x, 9)
	fmt.Printf("%p\n", &y) //变了
	fmt.Println(y)         //[0 0 0 0 0 0 0 0 0 0 9]
}

func addPtr(x *[]int) {
	*x = append(*x, 9)
}

func test3() {
	var a []byte
	b := make([]byte, 1)
	copy(b, a[0:1])
}

func test4() {
	s := make([]int, 2)
	mdSlice(s)
	fmt.Println(s)
}

func mdSlice(s []int) {
	fmt.Printf("%p\n", &s) //变了
	s[0] = 1
	s[1] = 2
}

func t5() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	x = Min(arr...)
	fmt.Printf("The minimum in the array arr is: %d", x)
}

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func t6() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 4, 3, 2, 1}
	op1(a, &b)
	fmt.Println(a, b) //[101 2 3 4 5] [105 4 3 2 1 200]
}

func op1(a []int, b *[]int) {
	a[0] += 100
	(*b)[0] += 100
	a = append(a, 200)
	*b = append(*b, 200)
}

func t7() {
	arr := make([]*bag7, 10)
	arr[0] = &bag7{}
	op2(arr)
	fmt.Println(arr[0].A) //100
}

type bag7 struct {
	A int
}

func op2(x []*bag7) {
	x = append([]*bag7{}, x...) //这一句可以不要
	x[0].A = 100
}
