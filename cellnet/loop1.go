package main

import (
	"fmt"
	"github.com/davyxu/cellnet/timer"
	"time"
)

//cellnet的timer包使用测试，等待timeour后输出第一条信息，一共输出10条
func main() {
	timeOutDur := time.Duration(3) * time.Second
	timer.NewLoop(nil, timeOutDur, func(loop *timer.Loop) {
		fmt.Printf("current time is %s\n", time.Now().Format("2006/01/02 15:04:05"))
	}, nil).Start()
	time.Sleep(31 * time.Second) //这里的用法和time.AfterFunc一样，主线程必须休眠
}
