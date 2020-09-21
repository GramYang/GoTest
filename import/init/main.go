package main

import (
	_ "GoTest/import/init/init"
	use "GoTest/import/init/use"
	"fmt"
)

//init执行一次
//this is Usage
//执行main
//多次导入一个包，这个包中的init函数只会执行一次
func main() {
	use.Usage()
	fmt.Println("执行main")
}
