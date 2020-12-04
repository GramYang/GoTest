package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	//测试对通道输入nil的结果：
	//如果通道的数据类型是int之类的，输入nil会编译不通过。
	//而如果数据类型是默认值为nil之类的，那就正常获取nil。
	//test1()
	//用通道传递切片，顺便说一句通道传递的是引用而不是值
	test2()
	//判断通道是否关闭
	//test3()
	//M个receivers，一个sender，sender通过关闭data channel说“不再发送”
	//test4()
	//一个receiver，N个sender，receiver通过关闭一个额外的signal channel说“请停止发送”
	//test5()
	//M个receiver，N个sender，它们当中任意一个通过通知一个moderator（仲裁者）关闭额外的signal channel来说“让我们结束游戏吧”
	//test6()
	//channel错误使用测试
	//t7()
}

func test1() {
	ch := make(chan error)
	go func(ch chan error) {
		ch <- errors.New("first")
		time.Sleep(time.Second * 2)
		ch <- errors.New("second")
		time.Sleep(time.Second * 2)
		ch <- nil
	}(ch)
	go func(ch chan error) {
		for {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println(i)
				} else {
					fmt.Println("ch已关闭")
				}
			case err := <-ch:
				fmt.Println(err)
			}
		}

	}(ch)
	time.Sleep(time.Second * 10)
}

func test2() {
	ch := make(chan []int, 3)
	b1 := []int{1, 2}
	b2 := []int{1, 2, 3}
	b3 := []int{1, 2, 3, 4}
	ch <- b1
	ch <- b2
	ch <- b3
	close(ch)
	fmt.Println(ch == nil) //false
	for c := range ch {
		fmt.Println(c)
	}
}

func test3() {
	c := make(chan int)
	fmt.Println(isClosed(c)) // false
	close(c)
	fmt.Println(isClosed(c)) // true
}

func isClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

//向关闭的通道发送会导致panic，返回true表示这个通道已经关闭并且panic了
func SafeSend(ch chan int, value int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

//二次关闭已关闭的通道会导致panic，返回false表示关闭失败因为该通道已经关闭了
func SafeClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}

//将通道和sync.Once一起封装，然后用once来关闭通道
type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan int)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

//也可以用mutex和bool与通道一起封装
type MyChannel2 struct {
	C      chan int
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel2() *MyChannel2 {
	return &MyChannel2{C: make(chan int)}
}

func (mc *MyChannel2) SafeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannel2) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}

func test4() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const MaxRandomNumber = 100000
	const NumReceivers = 100
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	dataCh := make(chan int, 100)
	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()
	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}
	wgReceivers.Wait()
}

func test5() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const MaxRandomNumber = 100000
	const NumSenders = 1000
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)
	dataCh := make(chan int, 100) //多个发送者一个接收者
	stopCh := make(chan struct{}) //一个发送者多个接收者
	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}
	// the receiver
	go func() {
		defer wgReceivers.Done()
		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()
	wgReceivers.Wait()
}

func test6() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	toStop := make(chan string, 1)
	var stoppedBy string
	// moderator，在toStop收到消息后关闭stopCh
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()
	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				//当value等于0时向toStop发送一个字符串然后退出，因为toStop可能已经关闭，因此设置了default:
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}
				select {
				case <-stopCh: //这里是stopCh的接收者，如果stopCh关闭则退出
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case dataCh <- value: //sender发送数据
				}
			}
		}(strconv.Itoa(i))
	}
	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				select {
				case <-stopCh: //这里是stopCh的接收者，如果stopCh关闭则退出
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case value := <-dataCh: //receiver接收数据
					if value == MaxRandomNumber-1 {
						select {
						case toStop <- "receiver#" + id: //向toStop发送数据后退出
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

func t7() {
	c := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
		}
		//close(c)//关闭c后，下面会无限输出0，说明读一个关闭的管道，会获得0值
		//注释了close(c)后会报死锁，因为当前只有一个协程，就是主协程（子协程已经结束），它阻塞了，自然就死锁了
		//必须把使用通道的两端都放入子协程中，才会避免死锁
	}()
	go func() {
		for {
			fmt.Print(<-c)
		}
	}()
	time.Sleep(1e9)
}
