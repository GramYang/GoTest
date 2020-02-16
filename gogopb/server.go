package main

import (
	chat "GoTest/gogopb/proto"
	"GoTest/gogopb/util"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/gogopb"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/davyxu/golog"
	"reflect"
)

var log = golog.New("server")

func main() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*chat.ChatReq)(nil)).Elem(),
		ID:    1024,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*chat.ChatAck)(nil)).Elem(),
		ID:    1023,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*chat.BytesReq)(nil)).Elem(),
		ID:    1022,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*chat.BytesAck)(nil)).Elem(),
		ID:    1021,
	})
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "server", "192.168.101.5:6789", queue)
	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
		case *cellnet.SessionClosed:
			log.Debugln("session closed: ", ev.Session().ID())
		case *chat.ChatReq:
			ack := &chat.ChatAck{
				Content: msg.Content + "nmsl",
				Status:  200,
			}
			ev.Session().Send(ack)
			//log.Infoln("client send: ", msg)
		case *chat.BytesReq:
			result := msg.Data
			if result == nil {
				log.Debugln("receive nil message")
			}
			origData, err := util.DESDecrypt(result, []byte("12345678"), []byte("87654321"))
			if err != nil {
				panic(err)
			}
			log.Debugln(string(origData))
		}
	})
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
