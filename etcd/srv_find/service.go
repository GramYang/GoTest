package srv_find

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

type ServiceInfo struct {
	IP string
}

type Service struct {
	Name    string
	Info    ServiceInfo
	stop    chan error
	leaseId clientv3.LeaseID
	client  *clientv3.Client
}

func NewService(name string, info ServiceInfo, endpoints []string) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Service{
		Name:   name,
		Info:   info,
		stop:   make(chan error),
		client: cli,
	}, err
}

func (s *Service) Start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for {
		select {
		case err := <-s.stop:
			_ = s.revoke()
			log.Printf("revoke the lease at time: %s\n", time.Now().String())
			return err
		case <-s.client.Ctx().Done():
			return errors.New("service closed")
		case ka, ok := <-ch:
			if !ok {
				log.Println("keep alive channel closed")
				_ = s.revoke()
				return nil
			} else {
				log.Printf("Recv reply from service: %s, ttl:%d", s.Name, ka.TTL)
			}
		}
	}
}

func (s *Service) Stop() {
	s.stop <- nil
}

func (s *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	info := &s.Info
	key := "services/" + s.Name
	value, _ := json.Marshal(info)
	res, err := s.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(res.ID))
	if err != nil {
		log.Fatal(err)
	}
	s.leaseId = res.ID
	return s.client.KeepAlive(context.TODO(), res.ID) //KeepAlive无限续租，通道不停的返回续租信息
}

func (s *Service) revoke() error {
	_, err := s.client.Revoke(context.TODO(), s.leaseId) //撤销续租了不用等ttl，直接删除
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("service:%s stop\n", s.Name)
	return err
}
