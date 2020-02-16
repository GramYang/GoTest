package main

import (
	"flag"
	"fmt"
)

func main() {
	//运行go run flag1.go -name 田所浩二 -age 24 -email 114514@163.com 你是一个 一个一个 一个啊~
	//输出
	//name: 田所浩二
	//age: 24
	//muscle: true
	//email: 114514@163.com
	//args: [你是一个 一个一个 一个啊~]

	//test1()

	//go run flag1.go -name 野兽先辈 -age 24 -student true -email 114514@1919.com -hello 好啊来啊 打在胸上 fa？！
	//输出
	//name: 野兽先辈
	//age: 24
	//vip: true
	//hello:
	//email: chenqionghe@sina.com
	//other: [true -email 114514@1919.com -hello 好啊来啊 打在胸上 fa？！]

	//test2()

	//运行go run flag1.go -test=12345
	//输出
	//test:default value
	//test1:<nil>
	//test:12345
	//test1:<nil>

	test3()
}

func test1() {
	namePtr := flag.String("name", "我是你哥哥", "姓名")
	agePtr := flag.Int("age", 24, "年龄")
	stuPtr := flag.Bool("学生", true, "是否是学生")

	var email string
	flag.StringVar(&email, "email", "1145141919@163.com", "邮箱")
	flag.Parse()
	args := flag.Args()
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("muscle:", *stuPtr)
	fmt.Println("email:", email)
	fmt.Println("args:", args)
}

type Value interface {
	String() string
	Set(string) error
}

type Hello string

func (p *Hello) Set(s string) error {
	v := fmt.Sprintf("Hello %v", s)
	*p = Hello(v)
	return nil
}

func (p *Hello) String() string {
	return fmt.Sprintf("%v", *p)
}

func test2() {
	namePtr := flag.String("name", "username", "姓名")
	agePtr := flag.Int("age", 18, "年龄")
	stuPtr := flag.Bool("student", true, "是否是学生")

	var email string
	flag.StringVar(&email, "email", "chenqionghe@sina.com", "邮箱")

	var hello Hello
	flag.Var(&hello, "hello", "hello参数")

	flag.Parse()
	others := flag.Args()
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("vip:", *stuPtr)
	fmt.Println("hello:", hello)
	fmt.Println("email:", email)
	fmt.Println("other:", others)
}

//定义一个全局变量的命令行接收参数
var _ = flag.String("test", "default value", "help message.")

//打印值的函数
func print1(f *flag.Flag) {
	if f != nil {
		fmt.Println(f.Value)
	} else {
		fmt.Println(nil)
	}
}

func test3() {
	//没有用flag.Parse()解析前
	fmt.Print("test:")
	print1(flag.Lookup("test"))
	fmt.Print("test1:")
	print1(flag.Lookup("test1"))

	//用flag.Parse()解析后
	flag.Parse()
	fmt.Print("test:")
	print1(flag.Lookup("test"))
	fmt.Print("test1:")
	print1(flag.Lookup("test1"))
}
