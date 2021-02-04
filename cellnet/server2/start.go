package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

type TestEchoACK struct {
	Msg   string
	Value int32
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    1234,
	})
}

const (
	TestAddress = "http://127.0.0.1:18802/echo"
)

func main() {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", TestAddress, queue)
	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			fmt.Println("server accepted")
		case *cellnet.SessionClosed:
			fmt.Println("session closed: ", ev.Session().ID())
		case *TestEchoACK:
			fmt.Printf("recv: %+v %v\n", msg, []byte("鲍勃"))
			val, exist := ev.Session().(cellnet.ContextSet).GetContext("request")
			if exist {
				if req, ok := val.(*http.Request); ok {
					raw, _ := json.Marshal(req.Header)
					fmt.Printf("origin request header: %s\n", string(raw))
				}
			}
			ev.Session().Send(&TestEchoACK{
				Msg:   "中文",
				Value: 1234,
			})
		}
	})
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
