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

type TestEchoACK struct {
	Msg   string
	Value int32
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestEchoACK")),
	})
}

const peerAddress = "127.0.0.1:17701"

func main() {
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			fmt.Println("server accepted")
		case *TestEchoACK:
			fmt.Printf("server recv %+v\n", msg)
			ev.Session().Send(&TestEchoACK{
				Msg:   msg.Msg,
				Value: msg.Value,
			})
		case *cellnet.SessionClosed:
			fmt.Println("session closed: ", ev.Session().ID())
		}

	})
	peerIns.Start()
	queue.StartLoop()
	queue.Wait() //一定要加这个，不然直接退出了
}
