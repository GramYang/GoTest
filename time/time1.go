package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

//time包常规操作
func main() {
	//打印当前时间，time.Unix测试，time.Time作为结构体域与json的互换
	test1()
	//一次定时器，有三种实现形式
	//test2()
	//循环定时器
	//test3()
	//Duration是指定的参数类型，由int64转换而来，time.Unix测试
	//test4()
	//判断time.Time的空值
	//t5()
	//time延时任务
	//t6()
}

type timebag struct {
	T time.Time
}

func test1() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nano := now.Nanosecond()
	date1 := time.Date(year, month, day, hour, minute, second, nano, time.Local)
	fmt.Println(now)                                    //2019-11-29 16:18:03.7328353 +0800 CST m=+0.001997701
	fmt.Println(year, month, day, hour, minute, second) //2019 November 29 16 18 3
	fmt.Println(date1)                                  //2019-11-29 16:18:03.7328353 +0800 CST
	fmt.Println(now.Format("2006-01-02 15:04:05"))      //2019-11-29 16:21:15
	//"2006-01-02 15:04:05"是固定写法，不是golang的诞生时间，只是简单的1234567
	fmt.Println(now.Format("2006/01/02 15:04:05"))               //2019/11/29 16:21:38
	fmt.Println(now.Format(".20060102"))                         //.20200630
	fmt.Println(now.Format(".20060102.150405"))                  //.2020630.150644
	fmt.Println(now.Nanosecond(), now.Unix())                    //508046200 1593942253
	fmt.Println(time.Unix(now.Unix(), 0))                        //2020-07-05 17:50:39 +0800 CST
	fmt.Println("time.Unix: ", time.Unix(1596108152461/1000, 0)) //time.Unix:  2020-07-30 19:22:32 +0800 CST
	t1 := "2019-01-08 13:50:30"
	timeTemplate := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate, t1, time.Local)
	fmt.Println("timestamp to second:", stamp.Unix()) //timestamp to second: 1546926630
	json1 := `{"createTime":1608624760}`
	var tb timebag
	_ = json.Unmarshal([]byte(json1), &tb)
	fmt.Println(tb) //{0001-01-01 00:00:00 +0000 UTC}出错了，不能直接转
	mapjson1 := map[string]interface{}{}
	_ = json.Unmarshal([]byte(json1), &mapjson1)
	mapjson1["createTime"] = time.Unix(int64(mapjson1["createTime"].(float64)/1000), 0)
	json2, _ := json.Marshal(&mapjson1)
	fmt.Println(string(json2)) //{"createTime":"1970-01-19T22:50:24+08:00"}
	var tb1 timebag
	_ = json.Unmarshal(json2, &tb1)
	fmt.Println(tb1) //{0001-01-01 00:00:00 +0000 UTC}，看来json下time.Time和int或者string的转换是行不通的
}

func test2() {
	/*
	   用sleep实现定时器
	*/
	fmt.Println(time.Now())
	time.Sleep(time.Second) //阻塞当前线程
	fmt.Println(time.Now())
	/*
	   用timer实现定时器，会阻塞当前线程
	*/
	timer := time.NewTimer(time.Second)
	fmt.Println(<-timer.C)
	/*
	   用after实现定时器，会阻塞当前线程
	*/
	fmt.Println(<-time.After(time.Second))
	//timer.Reset(time.Second) //重置
	//timer.Stop() //停用
}

func test3() {
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ticker.C)
	}
	ticker.Stop() //停止
}

func test4() {
	t := time.Duration(70) * time.Second
	t1 := t % time.Second
	t2 := t - t1
	fmt.Printf("%d %d", t1, t2) //0 70000000000
}

type timeWrap struct {
	A int
	B time.Time
}

func t5() {
	t := &timeWrap{}
	fmt.Println(t.A == 0)           //true
	fmt.Println(t.B == time.Time{}) //true
	var t1 timeWrap
	fmt.Println(t1.A == 0)           //true
	fmt.Println(t1.B == time.Time{}) //true
	if (t1.B == time.Time{}) {
		fmt.Println("equal") //equal
	}
	fmt.Println(reflect.ValueOf(t1.B).Interface() == reflect.ValueOf(time.Time{}).Interface()) //true
}

func t6() {
	//time.AfterFunc使用测试
	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(time.Second*2, func() {
		fmt.Println("123")
		wg.Done()
	})
	wg.Wait()
	//
	start := time.Now()
	end := <-time.After(1000 * 1000 * 1000 * 5)
	fmt.Println(start.Second(), end.Second()) //等5秒后输出59、4，可以看出秒数是60的余数
}
