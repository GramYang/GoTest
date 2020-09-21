package main

import (
	log "github.com/sirupsen/logrus"
)

//time="2020-06-30T11:02:58+08:00" level=info msg="A walrus appears" animal=walrus
//time="2020-06-30T11:06:27+08:00" level=info msg="A walrus appears without fields"
func main() {
	log.AddHook(&NHook{})
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	//不输出域
	log.Info("A walrus appears without fields")
}

type NHook struct{}

func (h *NHook) Fire(e *log.Entry) error {
	e.Data["114514"] = "1919810"
	return nil
}

func (h *NHook) Levels() []log.Level {
	return log.AllLevels
}
