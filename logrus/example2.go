package main

import (
	l "github.com/sirupsen/logrus"
	"os"
)

func init() {
	l.SetFormatter(&l.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.SetLevel(l.WarnLevel)
}

//输出
//{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2019-07-18T18:01:01+08:00"}
//{"level":"fatal","msg":"The ice breaks!","number":100,"omg":true,"time":"2019-07-18T18:01:01+08:00"}
func main() {
	l.WithFields(l.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	l.WithFields(l.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	l.WithFields(l.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	contextLogger := l.WithFields(l.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
