package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	//完整例子
	//test1()
	//查询用
	test2()
	//获取事件
	//go test3()
	//time.Sleep(10 * time.Second)
}

func test1() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	kv := clientv3.NewKV(cli)
	putRes, _ := kv.Put(context.TODO(), "/test/a", "something")
	fmt.Println("put: /test/a", putRes)
	_, _ = kv.Put(context.TODO(), "/test/b", "another")
	_, _ = kv.Put(context.TODO(), "/testxxx", "干扰")
	getRes, _ := kv.Get(context.TODO(), "/test/a")
	fmt.Println("get: /test/a", getRes)
	rangeRes, _ := kv.Get(context.TODO(), "/test/", clientv3.WithPrefix())
	fmt.Println(rangeRes)
	lease := clientv3.NewLease(cli)
	grantRes, _ := lease.Grant(context.TODO(), 10)
	_, _ = kv.Put(context.TODO(), "/test/expireme", "gone...", clientv3.WithLease(grantRes.ID))
	keepRes, _ := lease.KeepAliveOnce(context.TODO(), grantRes.ID)
	fmt.Println(keepRes)
	op1 := clientv3.OpPut("/hi", "hello", clientv3.WithPrevKV())
	opRes, _ := kv.Do(context.TODO(), op1)
	fmt.Println(opRes)
	txn := kv.Txn(context.TODO())
	txnRes, _ := txn.If(clientv3.Compare(clientv3.Value("/hi"), "=", "hello")).
		Then(clientv3.OpGet("/hi")).
		Else(clientv3.OpGet("/test/", clientv3.WithPrefix())).
		Commit()
	fmt.Println(txnRes)
	rangeWatch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
	for watchResp := range rangeWatch {
		for _, watchEvent := range watchResp.Events {
			fmt.Println(watchEvent.Type, string(watchEvent.Kv.Key))
		}
	}
	rangeWatch1 := cli.Watch(context.Background(), "", clientv3.WithPrefix())
	for watchResp := range rangeWatch1 {
		for _, watchEvent := range watchResp.Events {
			fmt.Printf("%s,%q,%q", watchEvent.Type, watchEvent.Kv.Key, watchEvent.Kv.Value)
		}
	}
}

func test2() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	kv := clientv3.NewKV(cli)
	key := "_svcdesc_hub"
	getRes, _ := kv.Get(context.TODO(), key, clientv3.WithPrefix())
	fmt.Printf("query key: %s value: %s\n", string(getRes.Kvs[0].Key), string(getRes.Kvs[0].Value))
}

func test3() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	rangeWatch := cli.Watch(context.TODO(), "_svcdesc_", clientv3.WithPrefix())
	for res := range rangeWatch {
		for _, ev := range res.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("put event, key: %s, value: %s", string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}
}
