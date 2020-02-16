package main

import (
	"GoTest/import/singleton/test1"
	"GoTest/import/singleton/test2"
)

//测试结果成功，这就是标准的单例写法，在包中用var包裹变量即可，根本不需要do.once
func main() {
	test1.Say()
	test2.Say()
}
