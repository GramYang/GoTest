package main

import (
	log "github.com/sirupsen/logrus"
)

//输出
//INFO[2019-07-19T09:12:25+08:00] Something noteworthy happened!
//WARN[2019-07-19T09:12:25+08:00] You should probably take a look at this.
//ERRO[2019-07-19T09:12:25+08:00] Something failed but I'm not quitting.
//FATA[2019-07-19T09:12:25+08:00] Bye.
func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true, //彩色log，只有TextFormatter有
		FullTimestamp: true,
	})
	log.SetLevel(log.TraceLevel) //默认为info等级，你可以设置为trace等级
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
