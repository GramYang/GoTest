package main

import (
	l "github.com/sirupsen/logrus"
)

//输出 time="2019-07-18T17:56:23+08:00" level=info msg="A walrus appears" animal=walrus
func main() {
	l.WithFields(l.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
