package main

import (
	"github.com/davyxu/golog"
	"os"
)

//测试cellnet的log包github.com/davyxu/golog
func main() {
	var log = golog.New("我是你爹")
	file, err := os.OpenFile("nmsl.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutptut(file)
	log.Debugln("我是你哥哥，我们两个都是你妈的儿子")
}
