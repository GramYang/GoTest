package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/cellnet/util"
	"reflect"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

const peerAddress = "127.0.0.1:17701"

type TestEchoACK struct {
	Msg   string
	Value int32
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestEchoACK")),
	})
}

func main() {
	done := make(chan struct{})
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Connector", "client", peerAddress, queue)
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		//连接成功后发送一个消息
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
			ev.Session().Send(&TestEchoACK{
				Msg:   "hello",
				Value: 1234,
			})
		//收到响应后就关闭
		case *TestEchoACK:
			fmt.Printf("client recv %+v\n", msg)
			done <- struct{}{}
		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})
	peerIns.Start()
	queue.StartLoop()
	<-done
}
