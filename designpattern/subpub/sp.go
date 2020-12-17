package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//发布订阅模式
func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)
	//eb是一个类实例，其中维护一个map，一个key对用一个通道切片，通道的数据里面也有key
	//eb.Subscribe向map中创建新键值对，该方法上锁
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)
	//publisTo中循环调用eb.Publish，其中上锁，向key对应的通道切片中发送数据
	go publisTo("topic1", "Hi topic 1")
	go publisTo("topic2", "Welcome to topic 2")
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, chans)
	}
	eb.rm.RUnlock()
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publisTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
